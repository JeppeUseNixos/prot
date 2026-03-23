{ self, ... }:
{
  perSystem =
    { pkgs, ... }:
    let
      version = "1.3.1";
      ldflags = [
        "-s"
        "-w"
        "-X github.com/z3co/prot/cmd.Version=${version}"
        "-X github.com/z3co/prot/cmd.Commit=${self.rev or "dirty"}"
      ];
    in
    {
      packages.dev = pkgs.buildGoModule {
        pname = "prot";
        inherit version;
        src = ./.;
        vendorHash = "sha256-1xQSQTUZUzykz8YXVnIp5bImU9cJCODiA3cWeb852w0=";
      };
			packages.default = pkgs.buildGoModule {
				pname = "prot";
				inherit version ldflags;
				src = pkgs.fetchFromGitHub {
					owner = "JeppeUseNixos";
					repo = "prot";
					tag = "v${version}";
					hash = "sha256-L/vY5/KzIB4R7JC+louQaTzLj5Axney/Vp71mZGXqdo=";
				};
				vendorHash = "sha256-1xQSQTUZUzykz8YXVnIp5bImU9cJCODiA3cWeb852w0=A";
			};
    };
}
