import { sleep, check } from 'k6';
import grpc from 'k6/net/grpc';

export const options = {
    vus: 2,
    duration: '3s',
};

const client = new grpc.Client();
client.load([], '../server/.protos/v1/board.proto');

export default async () => {
    client.connect('localhost:9095', {
        plaintext: true,
        timeout: '1s'
    });

    // CreateSubject *3 -> ListSubject --> defer DeleteSubject
    // CreateQuestion *5 -> ListQuestion -> Like / UnLike -> ListQuestion --> defer DeleteQuestion

    /* ListSubject
    const response = client.invoke('board.Board/ListSubject', {})
    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
    });

    console.log("ListSubject => "+response.body);
     */

    const newSubject = { title: 'gRPC'}
    const response = client.invoke('board.Board/CreateSubject', newSubject);

    console.log("CreateSubject => "+response.json);

    // const Subject = response.json().slideshow.slides[0];
    // check(Subject, {
    //     'title is gRPC': (s) => s.title === 'gRPC',
    //     'status is OK': (r) => r && r.status === grpc.StatusOK,
    // });

    console.log("Error: "+response.error)

    client.close();
    sleep(1);
};
