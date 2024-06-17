# fingermé

A golang based FINGERD(8) server for serving my my résumé as a .plan file.

## Background

Originally this was running in GKE, deployed with Helm, and fronted by a
load balancer. However, I don't really feel like spending ~10$ a month for
the `loadBalancerIP`.

## Usage

Query the server via a FINGER(1) client:

    $ task serve
    $ finger retr0h@localhost

Install dependencies

    $ task deps

Test packages:

    $ task unit

Test all:

    $ task test

List targets:

    $ task

## License

The [MIT][] License.

[MIT]: LICENSE
