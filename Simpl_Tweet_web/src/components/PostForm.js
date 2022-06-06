import React, {forwardRef}from 'react';
import {Button, Form, Input, Upload} from "antd";
import {UploadOutlined} from "@ant-design/icons";


const formItemLayout = {
    labelCol: {
        span: 6,
    },
    wrapperCol: {
        span: 14,
    },
};

export const PostForm = forwardRef((props, formRef) => {
    const normFile = (e) => {
        //console.log('Upload event:', e);

        if (Array.isArray(e)) {
            return e;
        }

        return e?.fileList;
    };

    //you need to transfer the state to Modal
    return (
        <Form
            name="validate_other"
            {...formItemLayout}
            ref={formRef}
        >
            <Form.Item
                label="Post"
                name="description"
                rules={[
                    {
                        required: true,
                        message: 'Please input your post',
                    },
                ]}>
                <Input.TextArea showCount maxLength={100} placeholder="Your post" />
            </Form.Item>

            <Form.Item
                name="upload"
                label="Upload"
                valuePropName="fileList"
                getValueFromEvent={normFile}
                extra="Pictures and Videos only and only 1 file each time"
            >
                <Upload name="logo" beforeUpload={ () => false} listType="picture" maxCount={1}>
                    <Button icon={<UploadOutlined />}>Click to upload</Button>
                </Upload>
            </Form.Item>

        </Form>
    );
}); //That is how to get the component to other component. FormRef is for the upper component a tool to get current component

