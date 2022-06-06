import React from "react";
import {Redirect, Route, Switch} from "react-router-dom";

import Login from"./Login"
import Register from"./Register"
import Home from"./Home"

//以下是一级路由
function Main(props) {
    const { isLoggedIn, handleLoggedIn} = props;


    const showLogin = () => {
        return isLoggedIn ? (
            <Redirect to="/home" />
        ) : (
            <Login handleLoggedIn={handleLoggedIn} />
        );
    };

    const showHome = () => {
        return isLoggedIn ? <Home /> : <Redirect to="/login" />;
    };


    //Route的好处是做到一对一的精确对应关系
    //exact match永远是用于可能会产生歧义的地方。比如说"/"鬼知道会match到哪儿，所以就用这个好了
    return(
        <div className="main">
            <Switch>
                <Route path="/" exact render={showLogin} />
                <Route path="/login" render={showLogin} />
                <Route path="/register" component={Register} />
                <Route path="/home" render ={showHome} />
            </Switch>
        </div>
    )
}

export default Main;