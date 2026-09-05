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
        debugpy

        numpy
        pympler
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
          echo "Python: $(python --version)"
          echo "ipykernel: $(python -c 'import ipykernel; print(ipykernel.__version__)')"
          echo "debugpy: $(python -c 'import debugpy; print(debugpy.__version__)')"
          echo "Run: jupyter notebook"
          echo "Happy Hacking 😎"
        '';
      };
    };
}