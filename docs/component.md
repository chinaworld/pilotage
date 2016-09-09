## Customized DevOps Component 

The Customized DevOps Component is the core concept in the ContainerOps, it's the part **DevOps Base Container**. Usually developer write DevOps plugins with *Bash*, *Python* and other script languages, but it has *runtime environment consistency* problem. When we got a script, we must setup a correct runtime environment to execute, it's a trivial and easy makes mistakes job. Now the container just fitly solved the *environment consistency* problem. Developer would use *Golang*, *Java*, *Rust* or any program languages or tools for the DevOps tasks, and share with runtime environment in contaienr image format. When we got the a container image, we don't care the environment at all and just care the function. 

### What's Customized DevOps Component

The customized DevOps component is a container image include DevOps programs or tools and runtime environment, it has specify inputs with environment parameters or REST API interface and specify outputs with callback interface.Now we use [Docker Registry Image specification](https://github.com/docker/distribution/blob/master/docs/spec/manifest-v2-2.md), and will support [ACI](https://github.com/appc/spec/blob/master/spec/aci.md) and [OCI Image specification](https://github.com/opencontainers/image-spec) soon.

We don't follow the [pods concept of Kubernetes](http://kubernetes.io/docs/user-guide/pods), but we execute the component in a pod of Kubernetes. We wanna to keep the component simply for sharing, it's very important for the community. If a DevOps task is complex which need more than one component, you should use multiple **Stages** and **Action**, and define the workflow in the Pilotage.

The component definition is open, we are glad to hear you feedback. Please open an [issue](https://github.com/containerops/pilotage/issues) to discuss.  

### How Write A DevOps Component

#### Write A Execute Task

#### How To Got And Export A Pipeline Event Data

##### Get A Event Data From REST API

##### Get A Event Data From Environment Variable

##### Export Event Data

#### How To Definition The Event Data With Editor

#### System Callback Environment Variable

When a new DevOps component container created, the DevOps workflow engine will set some environment variables automatically. All the variables is REST API URL, the DevOps component should call it for passing the component information or status to the DevOps workflow execute engine.

Sequence |  Variable       |  Value 
-------- | --------------- | ---------
   1     | COMPONENT_START | When the container of DevOps component start completely include all the dependencies started, the component should call the REST API of *COMPONENT_START* passing the start status status to the engine. Then the workflow execute engine will monitor the container status throw the orchestration tools like Kubernetes. And the execute engine will call the component passing the *event* data.
   2     | TASK_START      | When the container of DevOps component get all datas throw *event* data or volume data, call the REST API of *TASK_START* passing the task start execute status to the engine.
   3     | TASK_STATUS     | When the container of DevOps component execute the task, it should call the REST API of *TASK_STATUS* passing the interim outputs to the execute engine repeatly.
   4     | TASK_RESULT     | When the container of DevOps component execute successfully or failure, it should call the REST API of *TASK_RESULT* passing the result and final output to the execute engine. 
   5     | COMPONENT_STOP  | When the program of task in the container of DevOps component stop completely, it should call REST API of *COMPONENT_STOP* passing the stop status to the execute engine. The engine will notify the orchestration tools destory the container and release the resource. 

#### Build DevOps Component With Dockerfile And Acbuild

### How Registry Component In The Pilotage

#### Component Endpoint

#### Customized Environment Variables

#### Component Execute Timeout 

#### Volume Mount Location And Data

#### Component Execute Resource

#### Kubernetes Execute Configuration

### How Run Component In The Pilotage