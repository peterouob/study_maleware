#include <stdio.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>

int main()
{
    int server_socket;              // used for createing socket
    int client_socket;              // used for accepting connection
    struct sockaddr_in server_addr; // structure required by the `bind`
    char buffer[500] = "hello from server";
    char recv_buffer[500];

    server_socket = socket(AF_INET, SOCK_STREAM, 0); // create tcp socket

    server_addr.sin_family = AF_INET,                       // declaring the type
        server_addr.sin_port = htons(7001);                 // declaring the port
    server_addr.sin_addr.s_addr = inet_addr("192.168.3.1"); // declaring the host

    bind(server_socket, (struct server_addr *)&server_addr, sizeof(server_addr)); // Bind the server
    listen(server_socket, 5);

    client_socket = accept(server_socket, NULL, NULL); // accept incoming connections
    printf("Client connect! \n");
    send(client_socket, buffer, sizeof(buffer), 0);           // send data to the client
    recv(client_socket, recv_buffer, sizeof(recv_buffer), 0); // recv data from the client
    printf("clienct has sent :%s", recv_buffer);              // print out the data sent by client
    close(server_socket);                                     // close connection
    return 0;
}