import '../styles/App.css';
import React,{useState}from "react";

import TopBar from "./TopBar";
import Main from "./main";

import {TOKEN_KEY} from "../constants";

function App() {
    //isLoggedIn state = 前端local storage token有没有
    //需要后端返回token给我，所以useState接收后端的token response
    const [isLoggedIn, setIsLoggedIn] = useState(localStorage.getItem(TOKEN_KEY) ? true : false); //直接可以拿来用，这个等于window.localStorage
    //props就拥有了isLoggedIn了，但还是true/false
    //父传子：props。 子传父：callback函数
    //const [loginUser, setLoginUser] = useState("")

    const logout = () => { //直接把这个函数的接口传给Topbar.js，然后topbar就直接用这个函数来实现子传父
        console.log("logout")
        //step 1: delete token
        localStorage.removeItem(TOKEN_KEY);
        //step 2: use this to set state back to false
        setIsLoggedIn(false);
        localStorage.removeItem("username")
    };

    const loggedIn = (token, user) => {
        if (token) {
            localStorage.setItem(TOKEN_KEY, token);
            localStorage.setItem("username", user)
            //setLoginUser(user);
            setIsLoggedIn(true);

        }
    };

  return (
    <div className="App">
        <TopBar isLoggedIn = {isLoggedIn} handleLogout = {logout}/>
        <Main isLoggedIn = {isLoggedIn} handleLoggedIn = {loggedIn}/>
    </div>
  );
}

export default App;
