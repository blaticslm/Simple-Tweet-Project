import React, {useState} from "react"
import { Radio, Input } from 'antd';
import { SEARCH_TYPE } from "../constants";

const {Search} = Input

function SearchBar(props) {
    const [searchType, setSearchType] = useState(SEARCH_TYPE.all);
    const [error, setError] = useState(""); //error message

    const handleSearch = (input) => {
        if (searchType !== SEARCH_TYPE.all && input ==="") {
            setError("Please input your search keyword!");
            return;
        }
        setError("")
        //input search option and keywords from home
        props.handleSearch({type: searchType, keyword: input}) //second: this function takes the input as keyword
    }

    const searchTypeChange = (e) => {
        const search_Type = e.target.value //e.type.value may not be the same data type as const searchType
        setSearchType(search_Type); //First: define search type
        setError("")
        if(e.target.value === SEARCH_TYPE.all) {
            props.handleSearch({type: search_Type, keyword: ""})
        }

    };

    return(
        <div className="search-bar">
            <Search
                placeholder="input search text"
                enterButton="Search"
                size="large"
                onSearch={handleSearch}
                allowClear
                disabled={searchType === SEARCH_TYPE.all}
            />
            <p className="error-msg">{error}</p>

            <Radio.Group className="search-type-group" onChange={searchTypeChange} value={searchType}>
                <Radio value={SEARCH_TYPE.all}>View Gallery</Radio>
                <Radio value={SEARCH_TYPE.keyword}>keyword</Radio>
                <Radio value={SEARCH_TYPE.user}>User</Radio>
            </Radio.Group>

        </div>
    )
}

export default SearchBar