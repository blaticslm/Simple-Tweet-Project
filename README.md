# simple-tweet-project
This is the project for a simple tweeter-liked web application.

## Screenshot
<p align="center">
<img src="https://github.com/blaticslm/simple-tweet-project/blob/main/repo%20picture/QQ%E6%88%AA%E5%9B%BE20220607083659.png"  width="500">
<img src="https://github.com/blaticslm/simple-tweet-project/blob/main/repo%20picture/QQ%E6%88%AA%E5%9B%BE20220607083557.png"  width="500">
<img src="https://github.com/blaticslm/simple-tweet-project/blob/main/repo%20picture/QQ%E6%88%AA%E5%9B%BE20220607083541.png"  width="500">
</p>

## Demo
https://www.youtube.com/watch?v=Il3r-cv5hOI

This demo shows the following accomplishments : 
- Successfully registered by the new user
- Upload two types of post
- Users can only delete the post they created
- Search by keywords and user names 
- No need to refresh the web page to see the latest content update

## Program tools and environment 
#### Front end:
- Intellij
- Node.js
- React

#### Back end:
- Vscode
- Golang
- ElasticSearch database
- Google Compute Engine and Cloud Storage

## Front End Setup
- Node.js install and setup: 
https://docs.microsoft.com/en-us/windows/dev-environment/javascript/nodejs-on-windows

- Using command Terminal in IDE or CMD to create React project and test run: 
```
npx create-react-app {PROJECT_NAME_FOLDER}
cd {PROJECT_NAME_FOLDER}
npm start
```
## Back End Setup
#### Golang:
- Download and Install golang: https://go.dev/dl/
- Download golang extension in vscode: https://code.visualstudio.com/docs/languages/go

#### Google Cloud platform setup:
- VM instance firewall rules setup: https://geekflare.com/gcp-firewall-configuration/
- Create VM instance: https://www.youtube.com/watch?v=1FpMe8na64A and use the customized firewall rules into the VM
- Install golang inside the VM:
```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go

go version
```
#### SSH key to connect to VM instance using Vscode
- Using git bash (windows) to create ssh key
```
ssh-keygen -t rsa -f ~/.ssh/gcekey -C GMAIL_ACCOUNT
```
- Obtain the ssh key pair
```
cat ~/.ssh/gcekey.pub
```
- Go to VM instance to add the `SSH key` to the instance: https://medium.com/@vipiny35/how-to-add-ssh-keys-in-google-cloud-vm-instance-fa04d9cf7102
- Install SSH remote in vscode: https://code.visualstudio.com/docs/remote/ssh-tutorial
- Config connection:
```
ssh -i ~/.ssh/gcekey GMAIL_ACOUNT@GCE_EXTERNAL_IP
```
- Now follow the steps can access the VM machine from vscode now (At least I can do it since the detail steps are too long to go.)

## Backend structure
