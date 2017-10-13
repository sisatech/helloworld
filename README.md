# Hello World

Hello World is a simple webserver to use as an introduction to Vorteil. It hosts a webpage containing a simple hello world message with Vorteil's logo. The background colour of the page can be controlled by setting the BACKGROUND envrionment variable to any valid HTML colour code.

## Compilation

```sh
CGO_ENABLED=0 go build
```

To simplify the process of getting Hello World to run in a Virtual Machine, we set CGO_ENABLED=0 to direct Go to statically compile the binary without needing dynamically linked libraries.

## Running

The helloworld binary does not require any arguments.

```sh
./helloworld
```

Should return something like

```
2017/10/13 11:14:32 No background color set in BACKGROUND environment variable
2017/10/13 11:14:32 Binding port: 8888
```

As you might have guessed, you can set the background colour by specifying the BACKGROUND environment variable like so

```sh
BACKGROUND="0x000000" ./helloworld
```

You can connect to the web server by visiting http://localhost:8888/

## Vorteil

Testing helloworld in a Vorteil VM is easy.

```sh
vcli run
```

All configuration has already been applied within the helloworld.vcfg file.

Test the Vorteil app by connecting to http://localhost:8888/. Or use whichever port VCLI was able to bind. Once you've seen it working, compile the binary into a Vorteil Package to finish making the Vorteil app.

```sh
vcli package --icon vorteil.png
```

## VMS

Log in to VMS at https://go-vorteil.io

Upload the 'helloworld.vorteil' package you created by clicking on Applications on the sidebar and then the upload button in the upper-right corner of the main panel (it looks like an upward pointing blue arrow). When the popup window shows up, give your app a name like "helloworld".

Click on your new app in the file browser. This takes you to an application details page.

Click the "Actions" button, and then select "Deploy" from the drop-down menu.

Scroll down to the place where the number of VMs to deploy is listed, and change it to "1". Then click "Deploy".

VMS will redirect you to a deployment details page. Once the page finished loading, a url next to "URLs" will take you to your application's publically accessible internet webpage.
