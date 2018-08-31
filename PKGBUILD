# Maintainer: Trương Xuân Tính <xuantinh at gmail dot com>
pkgname=aur-pkg-status
pkgver=0.0.1
pkgrel=1
pkgdesc="A small utility to check the status of the AUR packages installed on your ArchLinux machine"
arch=('i686' 'x86_64')
license=('MIT')
url="https://github.com/tinhtruong/aur-pkg-status/"

source=("git+https://github.com/tinhtruong/aur-pkg-status.git")
md5sums=('SKIP')
depends=('pacman')
makedepends=('git' 'go')

build() {
	cd $srcdir/$pkgname
        export GOPATH=`pwd`:`pwd`/vendor
        cd src/github.com/tinhtruong/aur-pkg-status
	go build
}

package() {
	install -d "${pkgdir}/usr/bin"
        cp $srcdir/$pkgname/src/github.com/tinhtruong/aur-pkg-status/aur-pkg-status "${pkgdir}/usr/bin/aur-pkg-status"
}
