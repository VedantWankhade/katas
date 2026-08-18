{
  description = "Python Jupyter development environment";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = { nixpkgs, ... }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};

      python = pkgs.python313.withPackages (ps: with ps; [
        jupyter
        jupyterlab
        ipykernel

        numpy
        pandas
        matplotlib
      ]);
    in {
      devShells.${system}.default = pkgs.mkShell {
        packages = [
          python
        ];

        shellHook = ''
          echo "🐍 Python development environment"
          echo "Run: jupyter notebook"
          echo "Happy Hacking 😎"
        '';
      };
    };
}
