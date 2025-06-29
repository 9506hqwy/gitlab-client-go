#!/bin/bash
set -euo pipefail

BASE_DIR=$(dirname "$(dirname "$0")")
PKG_DIR="${BASE_DIR}/pkg/gitlab"
SOURCE_FILE="/tmp/client.go"

# Remove split files.
find "${BASE_DIR}/pkg/gitlab/" -type f -name "*.go" ! -name "*_test.go" -delete

# Output condition.
echo "const: $(grep -c '^const ' ${SOURCE_FILE})"
echo "type: $(grep -c '^type ' ${SOURCE_FILE})"
echo "func: $(grep -c '^func ' ${SOURCE_FILE})"

# Get categories from the client.go file.
CATEGORIES=$(grep -oP '^//.+ApiV4([A-Z][^A-Z. ]+)' "${SOURCE_FILE}" | sed -e 's/.*ApiV4\(.*\)/\1/' | sort | uniq)

# Split the client.go file into multiple files based on categories.
for CATEGORY in $CATEGORIES
do
    if [[ $CATEGORY == "Applications" || $CATEGORY == "Features" || $CATEGORY == "Groups" || $CATEGORY == "Jobs" || $CATEGORY == "Users" ]]; then
        continue
    fi

    FILE_NAME=${CATEGORY,,}

    # Write file header.
    cat > "${PKG_DIR}/${FILE_NAME}.go" <<EOF
package gitlab

import (
        "bytes"
        "context"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
        "net/url"
        "strings"
        "time"

        "github.com/oapi-codegen/runtime"
        openapi_types "github.com/oapi-codegen/runtime/types"
)

EOF

    {
        # Write const definiton.
        sed -n -Ee "/^\/\/.+ApiV4${CATEGORY}.+/ {
            N
            /const \($/!b
            :a
            N
            /\n\)$/!ba
            p
        }" "${SOURCE_FILE}"
        echo ""

        # Write type definition.
        sed -n -Ee "/^type .+ApiV4${CATEGORY}.+/ {
            /\{$/!bend
            :a
            N
            /\n\}$/!ba
            :end
            p
        }" "${SOURCE_FILE}"
        echo ""

        # Write func definition.
        sed -n -Ee "/^func .+ApiV4${CATEGORY}.+/ {
            :a
            N
            /\n\}$/!ba
            p
        }" "${SOURCE_FILE}"
        echo ""

    } >> "${PKG_DIR}/${FILE_NAME}.go"

done

# Write file header.
cat > "${PKG_DIR}/client.go" <<EOF
package gitlab

import (
        "bytes"
        "context"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
        "net/url"
        "strings"
        "time"

        "github.com/oapi-codegen/runtime"
        openapi_types "github.com/oapi-codegen/runtime/types"
)

EOF

# Write common code.
# shellcheck disable=SC2206
CATEGORIES=(${CATEGORIES})
OLD_IFS="$IFS"
IFS='|' COND_CATEGORY="${CATEGORIES[*]}"
IFS="$OLD_IFS"
{
    # Write type definition.
    sed -n -Ee "/^type .+/ {
        /(${COND_CATEGORY})/b
        /\{$/!bend
        :a
        N
        /\n\}$/!ba
        :end
        p
    }" "${SOURCE_FILE}"
    echo ""

    # Write func definition.
    sed -n -Ee "/^func .+/ {
        /(${COND_CATEGORY})/b
        :a
        N
        /\n\}$/!ba
        p
    }" "${SOURCE_FILE}"
    echo ""

} >> "${PKG_DIR}/client.go"


# Check condition.
echo "const: $(grep '^const ' -r "${PKG_DIR}" --exclude="*_test.go" | wc -l)"
echo "type: $(grep '^type ' -r "${PKG_DIR}" --exclude="*_test.go" | wc -l)"
echo "func: $(grep '^func ' -r "${PKG_DIR}" --exclude="*_test.go" | wc -l)"
