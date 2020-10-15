# go_rest_api

## api routes:
<ul>
    <li>for listing all articles, get request should be made at <b>/articles</b>. For pagination request format should be like this: <b>/articles?limit=20&after_id=10</b></li>
    <li>for posting a article, Post request should be sent at <b>/articles</b>with data in body should be in json format as below:
        {
            "Title": "",
            "SubTitle": "",
            "Content": ""
        }
    </li>
    <li>for a single article, get request should be sent on <b>/articles/article_id</b></li>
    <li>for searching a article query should be sent at <b>/articles/search?q=query_to_be_sent</b></li>

</ul>