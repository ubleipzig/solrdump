Summary:    Dump SOLR fields efficiently with cursors.
Name:       solrdump
Version:    0.1.7
Release:    0
License:    GPL
BuildArch:  x86_64
BuildRoot:  %{_tmppath}/%{name}-build
Group:      System/Base
Vendor:     Leipzig University Library, https://www.ub.uni-leipzig.de
URL:        https://github.com/ubleipzig/solrdump

%description

Dump SOLR fields efficiently with cursors.

%prep

%build

%pre

%install
mkdir -p $RPM_BUILD_ROOT/usr/local/sbin
install -m 755 solrdump $RPM_BUILD_ROOT/usr/local/sbin

%post

%clean
rm -rf $RPM_BUILD_ROOT
rm -rf %{_tmppath}/%{name}
rm -rf %{_topdir}/BUILD/%{name}

%files
%defattr(-,root,root)

/usr/local/sbin/solrdump

%changelog
* Thu Dec 6 2016 Martin Czygan
- 0.1.2 initial release
