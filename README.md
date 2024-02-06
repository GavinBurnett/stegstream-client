Overview

Stegstream is made up of a server and client program. The server program uses steganography to hide a file inside a MP3 music file. The music file is then made available via the server program streaming it, and it can be listened to using a web browser in the same way as normal streaming. Running the client program will extract the hidden file from the music file streamed by the server. This method allows files to be transferred in an invisible way, as any observer will simply see a music streaming service.

Design Goal

To provide a practical application of steganography that is easy for non technical people to use in order to privately distribute files to a large number of people.

Installation

Download the archive file containing the client executable for the relevant operating system from the releases. Validate the PGP signature if an integrity check is advisable, and extract the executable from the archive file. Place the executable in a suitable directory; E.G /home/username/stegstream

GnuPG Signing Key: http://pgp.mit.edu/pks/lookup?op=get&search=0x203092F792253A6F

Getting the hidden file from the server

Open a command prompt, and cd into the directory containing the client program.

Enter the following:

./stegstream-client server URL

An example:

./stegstream-client http://localhost:8080/Audio

This will download the streamed music file and extract the hidden file from it. The hidden file will be written to the same directory that the client program is stored in.

Things to consider

The hidden file is left on disk after the client program has finished. It is the responsibility of the user to delete this file after reading if secrecy is important.

In order for the client to see the server, make sure the computer running the server is network visible from the computer running the client – use ping to check visibility.
