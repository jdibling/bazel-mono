  # Bazel Monorepo experimentation project

  bazel-mono is a simple project to experiement with Bazel and modern SaaS technologies like Lambda and Pulumi, architected using a microservices approach.

  The main objective of this repo is to fugure out how to use the Bazel build system.

  We will build a fairly simple sample project: an audio file conversion website. Our website will have a space where you can drop media files such as MP3 or MP4, upload those files to our backend service, and the file will be transcoded to FLAC, where you can then download it to your laptop. We will use the following tools and services:

  * Go will be the language of choice for backend serverless functions
  * We will use a serverless architecture for the backend
  * AWS will be the cloud provider
  * The frontend will be an Angular website, hosted on AWS
  * Feature flags will be implemented using SSM and will control the transcode process in some way
  * Pulumi will be used for our IaC tool
  * DynamoDB will be our database

Here's how the system will work.

* The end user will navigate to the service's website, which will be a static S3 website
* The user will drag-and-drop a media file in a specific location on the website. 
* The user will click an upload button, and then  the following sequence of events will take place:
  * The frontend will call `POST /xfer/` 
  * The backend will create generate a new `SessionId`
  * The backend will create a presigned S3 URL pointing to a bucket
  * The backend will create add new `Session` record to the database, including the SessionID, a bucket name and key where the media file will be uploaded, and a session status of `NotUploaded`
  * The backend will return the SessionId and the S3 URL back to the frontend
  * The frontend will upload the media file to the presigned URL
  * The frontend will add a `Transcode` button to the frontend
* Once the file has been uploaded, an S3 trigger will fire, and set the session status to `ReadyToXcode`
* The user will click the transcode button, and thje following happens:
  * The frontend sends `PUT /xcode/{sessionId}`
  * If the session is not in state `ReadyToXcode` a `400 Wrong State` error is returned.
  * If the session is in `ReadyToXcode` state, the transcoding job is started.
* 



