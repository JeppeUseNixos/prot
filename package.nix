{ self, ... }:
{
  perSystem =
    { pkgs, ... }:
    let
      version = "1.0.1";
      ldflags = [
        "-s"
        "-w"
        "-X github.com/z3co/prot/cmd.Version=${version}"
        "-X github.com/z3co/prot/cmd.Commit=${self.rev or "dirty"}"
      ];
    in
    {
      packages.default = pkgs.buildGoModule {
        pname = "prot";
        inherit version ldflags;
        src = ./.;
        vendorHash = "sha256-1xQSQTUZUzykz8YXVnIp5bImU9cJCODiA3cWeb852w0=";
      };
    };
}
