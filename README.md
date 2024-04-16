## factorkube

Implementing factorio game with Kubernetes for learning purposes.

Here is an example of UI:

![image](https://github.com/overmesgit/factorio/assets/2367946/e752fea3-63bd-4f9f-912c-e7e07287c3d5)

You can see a map with cells and available buildings on the right side.
Users can click on the building and put it on the map. 
Each building is a container that runs inside Kubernetes.
Buildings can communicate through GRPC API to exchange resources.

The belt is the main structure to move resources between buildings.
A manipulator is needed to move items from buildings like furnaces and manufacture.
In current implementation two mines are available: iron and coal.
They can be used in furnaces to produce iron plates.
Iron plates can be used to produce gear.

Here is a full video with the game process:

https://github.com/overmesgit/factorio/assets/2367946/e137274a-0a67-4f28-a243-fbe3dc6bf26d

This is what it looks like in Kubernetes.
This map:

![image](https://github.com/overmesgit/factorio/assets/2367946/9e3b12d0-d9ad-42db-8770-f264b081cccf)

Has this pods:

![image](https://github.com/overmesgit/factorio/assets/2367946/8a4bce67-68dc-4129-81ea-ece686c824c5)
