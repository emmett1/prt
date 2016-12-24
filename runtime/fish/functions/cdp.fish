function cdp
    set portdir (cat /etc/prt/config.toml | string match -r 'portdir.*' | string trim -c '" ')
    set loc (prt loc $argv ^/dev/null)

    if test "$portdir/$loc"
        cd $portdir/$loc
    else
        cd $portdir
    end
end
