%define debug_package   %{nil}
%define _build_id_links none
%define _name metabypass
%define _prefix /opt
%define _version 1.01.00
%define _rel 0
%define _arch x86_64
%define _binaryname screwUmeta

Name:       metabypass
Version:    %{_version}
Release:    %{_rel}
Summary:    News medias publisher on Meta

Group:      misc
License:    GPL2.0
URL:        https://github.com/jeanfrancoisgratton/metabypass

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
News medias publisher on Meta

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Sat Oct 21 2023 RPM Builder <builder@famillegratton.net> 1.01.00-0
- Disabled CSS (jean-francois@famillegratton.net)
- Added info on go generate (jean-francois@famillegratton.net)
- Changed package name (builder@famillegratton.net)
- APK building script fixes (builder@famillegratton.net)
- Added newline to Debian's control file (builder@famillegratton.net)
- Permission fixes (builder@famillegratton.net)

* Sat Oct 21 2023 RPM Builder <builder@famillegratton.net> 1.00.00-0
- new package built with tito

