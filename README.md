# How to use benthos in knative

Benthos allow you to create custom plugins that extend benthos components.
In this example, we have created a new function in bloblang processor named len_s.
This dummy function only calculate the len of the string.

- **Prerequisites**

    - Knative in kubernetes

Once you have knative installed you only need to build the image in your local:
```docker build . -t dev.local/benthos-knative-plugin:latest```

We need to define the build with the prefix dev.local because knative skip docker hub if the image start with it.

Then, run this command:

```kubectl apply -f service.yaml```

Check that the service is up and running:

```kn service list```

Pick the url for the service benthos-knative and run the next curl:

Windows: ```curl -X POST -H "Content-type: application/json" --data "{\"foo\": \"I love Benthos!\"}" http://benthos-knative.default.{YOUR_IP}.sslip.io/post```

Linux: ```curl -v POST -d '{"foo": "I love Benthos!"}' http
://benthos-knative.default.{YOUR_IP}.sslip.io/post```