import React, {useEffect, useState} from "react"
import SearchBar from "../components/SearchBar";
import PhotoGallery from "../components/PhotoGallery";
import CreatePostButton from "./CreatePostButton";

import axios from "axios";
import {BASE_URL, SEARCH_TYPE, TOKEN_KEY} from "../constants";
import {Button, Col, message, Row, Spin, Tabs} from "antd";
import {DeleteOutlined, LoadingOutlined} from "@ant-design/icons";

const { TabPane } = Tabs;

const antIcon = (
    <LoadingOutlined
        style={{
            fontSize: 24,
        }}
        spin
    />
);

function Home(props) {

    const [posts, setPosts] = useState([]);
    const [activeTab, setActiveTab] = useState("image");
    const [videoArray, setVideoArray] = useState([]);
    const [imageArray, setImageArray] = useState([]);
    const [searchOption, setSearchOption] = useState({
        type: SEARCH_TYPE.all,
        keyword: ""
    });

    const [isLoading, setisLoading] = useState(false);

    //when enter the page first time, it renders. ---> did mount
    //when there is data change, it renders. ---》 didupdate
    //update: post
    useEffect(() => {
        setisLoading(true);
        fetchPos(searchOption)

    }, [searchOption]);

    useEffect(() => {
        setisLoading(true);
        fetchPos(searchOption)

    }, [activeTab])

    //render rules: post --> image --> PhotoGallery
    //              post --> video --> Home

    //new post is coming --> videoArray gets all videos from post

    //deleting the video --> directly delete from the local videoArray --> faster than refetch, so no synchronization problems!

    //once the new post is coming, then the page will be rendered
    useEffect(() => {
        setVideoArray(posts != null? posts.filter((item) => item.type === "video") : [])
        setImageArray(posts != null? posts.filter((item) => item.type === "image") : [])
     },[posts])

    //　When each time fetching, it will get all the videos and images from the server
    const fetchPos = (option) => {
        const {type, keyword} = option;
        let url; //scope is only existing in this function

        if(type === SEARCH_TYPE.all) {
            url = `${BASE_URL}/search`
        } else if(type === SEARCH_TYPE.user) {
            url = `${BASE_URL}/search?user=${keyword}`
        } else {
            url = `${BASE_URL}/search?keywords=${keyword}`
        }

        const opt = {
            method: 'GET',
            url: url,
            headers: {
                Authorization: `Bearer ${localStorage.getItem(TOKEN_KEY)}`
            }
        };

        axios(opt)
            .then((res) => {
                if(res.status === 200) {
                    setPosts(res.data);
                    setisLoading(false)
                }
            })
            .catch((err) => {
                setisLoading(false)
                console.log("Fetch Fail! ", err.message);
                message.error("Fetch failed!");
            });
    }

    const videoDelete = (post_id) => {
        if (window.confirm(`Are you sure you want to delete this VIDEO Post?`)) {

            //get rest of videos into array
            const newVideo_list = videoArray.filter((item) => item.id !== post_id)

            const option = {
                method: "DELETE",
                url: `${BASE_URL}/post/${post_id}`,
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem(TOKEN_KEY)}`
                }
            };

            axios(option)
                .then((res) => {
                    if(res.status === 200) {
                        setVideoArray(newVideo_list)
                    }
                })
                .catch((err) => {
                    console.log("Delete Fail! ", err.message)
                    message.error("Delete failed!");
                });
        }
    }

    const renderPosts = (type) => {

        if(type === "image") {
            if(!imageArray || imageArray.length === 0) {
                return <div>No data!</div>;
            }

            //obtaining the image type post
            const imageArr = imageArray
                //.filter((item) => item.type === "image")
                .map((item) => {
                    return {
                        postId: item.id,
                        src: item.url,
                        user: item.user,
                        caption: item.message,
                        thumbnail: item.url,
                        thumbnailWidth: 300,
                        thumbnailHeight: 200
                    };
                });
            return <PhotoGallery images = {imageArr} isLoading={isLoading}/>

        } else if(type === "video") {
            //gutter distance is between two windows
            //col span is in Antd web. Full=24, 2 is 12, 3 is 8, 4 is 6
            if(!videoArray || videoArray.length === 0) {
                 return <div>No data!</div>;
            }

            return(
                isLoading ?
                    <div className="loading_spin">
                        <Spin tip ="I am running!" size = "large" indicator={antIcon}/>
                    </div>
                    :
                    <Row gutter={32}>
                        {videoArray
                            //.filter((item) => item.type === "video")
                            .map((post) => (
                                <Col span={6} key={post.url}>
                                    <div className="videoDelete">
                                        <Button
                                            className="deleteButton"
                                            key="deletePost"
                                            type="primary"
                                            icon={<DeleteOutlined />}
                                            size="small"
                                            onClick={function(){videoDelete(post.id)}}
                                            hidden={localStorage.getItem("username") !== post.user}
                                        >Delete Post</Button>
                                    </div>
                                    <video width="250" src={post.url} controls={true} className="video-block" />
                                    <p>{post.user}: {post.message}</p>
                                </Col>)
                            )
                        }
                    </Row>
            );

        }
    };

    const handleSearch = (option) => {
        //console.log(option)
        const {type, keyword} = option;
        //update search option
        setSearchOption({type: type, keyword: keyword})
    }

    const showPost = (postType) => {
        setActiveTab(postType);
        setTimeout( () => {
            fetchPos({type: SEARCH_TYPE.all, keyword: ""})
        }, 3000); //giving 3 seconds for uploading and fetching
    }

    const upload = <CreatePostButton onShowPost={showPost}/>

    return(
        // tabBarExtraContent add extra upload button(in this case) to the tab bar
        //TabPane is the content under the tab
        <div className="home">
            <SearchBar handleSearch={handleSearch}/>
            <div className="display">
                <Tabs
                    onChange={(key) => setActiveTab(key)}
                    defaultActiveKey="image"
                    activeKey={activeTab}
                    tabBarExtraContent={upload}
                >
                    <TabPane tab="Images" key="image">
                        {renderPosts("image")}
                    </TabPane>
                    <TabPane tab="Videos" key="video">
                        {renderPosts("video")}
                    </TabPane>
                </Tabs>

            </div>
        </div>
    )
}

export default Home