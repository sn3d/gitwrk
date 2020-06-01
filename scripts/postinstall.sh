# We need to manually change the chmod because go-releaser
# publish binary with 666.
if [ "$1" = "configure" ]; then
    chmod 755 /usr/local/bin/gitwrk
fi