import React from "react"

import { Form, Input, Button, Checkbox, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import axios from "axios";
import { Link } from "react-router-dom";
import {BASE_URL} from "../constants";

function Login(props) {

    const {handleLoggedIn} = props

    const onFinish = (values) => {
        console.log('Received values of form: ', values);
        //要用这个发送给后端了，然后得到的token传给父
        //case 1: 得到token --> 子传父
        //case 2: display error
        const{username, password} = values

        const option = {
            method: 'POST',
            url: `${BASE_URL}/signin`,
            data: {
                username: username,
                password: password
            },
            headers: { "Content-Type": "application/json" }
        }; //得到信息的选项

        axios(option)
            .then((res) => {
                console.log(res.data);
                if(res.status === 200) {
                    const {data} = res; //see response
                    //接下来开始子传父
                    handleLoggedIn(data, username);
                    message.success("Login succeed! ");

                }
            })
            .catch((err) => {
                console.log("Login Fail! ", err.message)
                message.error("Login failed!");
            });


    };

    return(
        <Form
            name="normal_login"
            className="login-form"
            onFinish={onFinish}
        >
            <Form.Item
                name="username"
                rules={[
                    {
                        required: true,
                        message: 'Please input your Username!',
                    },
                ]}
            >
                <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
            </Form.Item>
            <Form.Item
                name="password"
                rules={[
                    {
                        required: true,
                        message: 'Please input your Password!',
                    },
                ]}
            >
                <Input
                    prefix={<LockOutlined className="site-form-item-icon" />}
                    type="password"
                    placeholder="Password"
                />
            </Form.Item>

            <Form.Item>
                <Button type="primary" htmlType="submit" className="login-form-button">
                    Log in
                </Button>
                Or <Link to="/register"> register now!</Link>
            </Form.Item>
        </Form>
    )

}

export default Login