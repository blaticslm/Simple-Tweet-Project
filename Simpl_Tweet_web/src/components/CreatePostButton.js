import React, {Component} from 'react';
import {Button, message, Modal} from "antd";
import {PostForm} from "./PostForm";
import {BASE_URL, TOKEN_KEY} from "../constants";
import axios from "axios";

class CreatePostButton extends Component {
    state = {
        popUp: false,
        loading: false
    }

    showPopUp = () => {
        this.setState({
            popUp: true
        });
    };

    handleCancel = () => {
        this.setState({
            popUp: false
        });
    }

    handleOk = () => {
        this.setState({
            loading: true,
        });
        this.postForm
            .validateFields()
            .then( form => {
                //step 1: get post info
                //step 2: send post to server
                //console.log(form);
                const{ description, upload } = form; //from console log: 2 parts from the form
                const{ type, originFileObj} = upload[0]; //from upload array: type and originalFileObj
                const postType = type.match(/^(image|video)/g)[0];

                if(postType) {
                    let formData = new FormData();
                    formData.append("message", description);
                    formData.append("media_file", originFileObj);

                    const option = {
                        method: 'POST',
                        url: `${BASE_URL}/upload`,
                        data: formData,
                        headers: {
                            "Authorization": `Bearer ${localStorage.getItem(TOKEN_KEY)}`
                        }
                    }; //得到信息的选项

                    axios(option)
                        .then((res) => {
                            if(res.status === 200) {
                                //Step 1: clear form content
                                //Step 2: close Modal
                                //Step 3: confirmLoading

                                message.success("Upload successfully!");
                                this.postForm.resetFields();
                                this.handleCancel();
                                this.props.onShowPost(postType)
                                this.setState({loading: false});

                            }
                        })
                        .catch((err) => {
                            console.log("Upload Fail! ", err.message)
                            message.error("Upload failed!");
                        });
                }

            })
            .catch(() => {
                message.error("You don't have enough contents!")
            })


    }

    render() {
        const { popUp, loading} = this.state
        return (
            <div>
                <Button type="primary" onClick={this.showPopUp}>
                    Create new post
                </Button>
                <Modal
                    title="Create new post"
                    visible={popUp}
                    onOk={this.handleOk} //update
                    onCancel={this.handleCancel}
                    confirmLoading={loading}
                    footer={[
                        <Button key="back" onClick={this.handleCancel}>
                            Return
                        </Button>,
                        <Button key="submit" type="primary" onClick={this.handleOk}>
                            Create
                        </Button>,
                    ]}>
                    <PostForm ref={ (postFormInstance) => {this.postForm = postFormInstance;}}/>
                </Modal>
            </div>

        );
    }
}

export default CreatePostButton;