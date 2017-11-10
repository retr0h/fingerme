# containumé

A container consisting of a golang FINGERD(8) server serving
my résumé as a .plan file.

## Usage

Building the container:

    $ make build

Running the container locally:

    $ make run

Query the server via a FINGER(1) client:

    $ pacaur -S netkit-bsd-finger
    $ finger retr0h@localhost  
    [localhost.localdomain]
    User: retr0h
    Plan: Testing testing 123

## TODO

* Reimplement the fingerd server to better learn golang.
* Create the fingerd user with known UID in Dockerfile.

## License

MIT
