import React from "react";
import logo from "../assets/images/logo.svg";

import{LogoutOutlined} from "@ant-design/icons";

function TopBar(props) {
    const{ isLoggedIn , handleLogout} = props; //因为传进来是个实例，还不是

    return(
        <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <span className="App-title">Simple Tweet</span>
            {
                isLoggedIn ? <LogoutOutlined className="logout" onClick={handleLogout} /> : null //ture就有button，false就没有

            }
        </header>
    );
}

export default TopBar;
