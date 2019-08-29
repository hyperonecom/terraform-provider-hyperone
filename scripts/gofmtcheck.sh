#!/usr/bin/env bash
# Source: https://github.com/terraform-providers/terraform-provider-template/blob/748304bc4709326d5cf4642adf3b0cb09b91a277/scripts/gofmtcheck.sh
# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(gofmt -l `find . -name '*.go' | grep -v vendor`)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit 1
fi

exit 0