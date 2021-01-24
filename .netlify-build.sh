curl -LO "https://github.com/bazelbuild/bazelisk/releases/download/v1.1.0/bazelisk-linux-amd64"
mkdir -p bin
mv bazelisk-linux-amd64 bin/bazel
chmod +x bin/bazel
bin/bazel --batch build  //draftcode.osak.jp:final_layout
