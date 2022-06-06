import React, {useEffect, useState} from 'react';
import Gallery from "react-grid-gallery";
import PropTypes from "prop-types";
import {Button, message, Spin} from "antd";
import {DeleteOutlined, LoadingOutlined} from "@ant-design/icons";
import axios from "axios";

import { BASE_URL, TOKEN_KEY } from "../constants";

//https://github.com/benhowell/react-grid-gallery/blob/master/examples/demo4.js
const captionStyle = {
    backgroundColor: "rgba(0, 0, 0, 0.6)",
    maxHeight: "240px",
    overflow: "hidden",
    position: "absolute",
    bottom: "0",
    width: "100%",
    color: "white",
    padding: "2px",
    fontSize: "90%"
};

const wrapperStyle = {
    display: "block",
    minHeight: "1px",
    width: "100%",
    border: "1px solid #ddd",
    overflow: "auto"
};

const antIcon = (
    <LoadingOutlined
        style={{
            fontSize: 24,
        }}
        spin
    />
);

function PhotoGallery(props) {
    const [imagesArr, setImagesArr] = useState(props.images);
    const [curImageIndex, setCurImageIndex] = useState(0);
    const [postUser, setPostUser] = useState("");

    const imageArr = imagesArr.map((image) => {
        //customOverlay: something above the image, react-grid-gallery
        return {
            ...image,
            customOverlay:(
                <div style={captionStyle}>
                    <div>{`${image.user}: ${image.caption}`}</div>
                </div>
            )
        };
    });


    useEffect(() => {
        setImagesArr(props.images);
    }, [props.images]);

    const onDeletePost = () => {

        if (window.confirm(`Are you sure you want to delete this Post?`)) {
            //step 1: get image with post to be deleted
            //step 2: remove the image from image array
            //step 3: send delete request

            const currentImage = imagesArr[curImageIndex]
            const newImageArr = imagesArr.filter((image, index) => index !== curImageIndex) //filtering out the image that does not have designate index

            const option = {
                method: "DELETE",
                url: `${BASE_URL}/post/${currentImage.postId}`,
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem(TOKEN_KEY)}`
                }
            };

            axios(option)
                .then((res) => {
                    if(res.status === 200) {
                        console.log("delete picture successfully");
                        setImagesArr(newImageArr);
                    }
                })
                .catch((err) => {
                    console.log("Delete Fail! ", err.message);
                    message.error("Delete failed!");
                });
        }
    }


    const onCurrentImageChange = (index) => {
        console.log(index)
        setCurImageIndex(index); //当前点击的图片的index
        try{
            setPostUser(imagesArr[index].user);
        } catch (e) {
            setPostUser("");
        }

    }

    return(
        props.isLoading ?
            <div className="loading_spin">
                <Spin tip ="I am running!" size = "large" indicator={antIcon} />
            </div>
            :
            <div style={wrapperStyle}>
                <Gallery
                    images={imageArr}
                    enableImageSelection={false}
                    backdropClosesModal={true}
                    currentImageWillChange={onCurrentImageChange}
                    customControls={[
                        <Button
                            key="delete Post"
                            type="primary"
                            icon={<DeleteOutlined />}
                            size="small"
                            onClick={onDeletePost}
                            hidden={localStorage.getItem("username") !== postUser}
                        >Delete Post</Button>
                    ]}
                />
            </div>
    );
}

//prop's types, so the inside props should be like this
//react has this thing
PhotoGallery.propTypes = {
    images: PropTypes.arrayOf(
        //every element's check
        PropTypes.shape({
            user: PropTypes.string.isRequired,
            caption: PropTypes.string.isRequired,
            src: PropTypes.string.isRequired,
            thumbnail: PropTypes.string.isRequired,
            thumbnailWidth: PropTypes.number.isRequired,
            thumbnailHeight: PropTypes.number.isRequired
        })
    ).isRequired //array is required
};

export default PhotoGallery;