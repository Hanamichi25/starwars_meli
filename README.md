# Starwars

Este proyecto contiene la solución al problema propuesto

## Pre-requisitos.

- Se debe contar con una cuenta AWS.
- Se debe crear un bucket de s3, donde se alojará el build de la aplicación

## Installation

- Encontrará un script _build.sh_ El cual se encarga de generar el ejecutable de la aplicación, en este debe modificar el nombre del bucket en la linea _8_ del script

Para realizar la ejecución:

```bash
./build.sh
```

## Infra

Debe tener en cuenta el archivo _all_vars.tfvars_, cambiando los valores a los que corresponda su aplicación.

```terra
terraform init
terraform plan -var-file=all_vars.tfvars
terraform apply -var-file=all_vars.tfvars
```

Una vez ejecutado encontrará el LB al cual podrá ejecutar las peticiones necesarias en el servicio.

## Diagrama arquitectura

https://github.com/Hanamichi25/starwars_melitree/main/data/arquitectura.jpg

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
