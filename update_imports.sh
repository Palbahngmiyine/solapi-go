#!/bin/bash

# List of files to update
files=(
  "examples/test/smsTest.go"
  "examples/Group/getGroupMessageLIst/getGroupMessageList.go"
  "examples/Group/createGroup/createGroup.go"
  "examples/Group/getGroupList/getGroupList.go"
  "examples/Group/deleteGroup/deleteGroup.go"
  "examples/Group/getGroup/getGroup.go"
  "examples/Group/addGroupMessage/addGroupMessage.go"
  "examples/Group/sendGroup/sendGroup.go"
  "examples/Message/getMessageLIst/getMessageList.go"
  "examples/Stroage/getFileList/getFileList.go"
  "examples/Stroage/upload/uploadFile.go"
)

# Update import paths in each file
for file in "${files[@]}"; do
  if [ -f "$file" ]; then
    sed -i 's|"github.com/solapi/solapi-go/v2/pkg/solapi"|"github.com/solapi/solapi-go/pkg/solapi"|g' "$file"
    echo "Updated $file"
  else
    echo "File not found: $file"
  fi
done

echo "Import paths updated in all example files."
