import { useState, useEffect } from 'react'
import axios from "@/axios/axios";

type Profile = {
    Email: string,
    Id: string,
}

const defaultProfile: Profile = {Email: "", Id: ""}


function Page() {
    const [data, setData] = useState<Profile>(() => defaultProfile)
    const [isLoading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true)
        axios.get('/users/profile')
            .then((res) => res.data)
            .then((data) => {
                setData(data)
                setLoading(false)
            })
    }, [])

    if (isLoading) return <p>Loading...</p>
    if (!data) return <p>No profile data</p>

    return (
        <div>
            <p>用户信息：{JSON.stringify(data)}</p>
            <br></br>
            <p>邮箱: {data.Email}</p>
            <p>用户id: {data.Id}</p>
            {/*<p>{data.bio}</p>*/}
        </div>
    )
}

export default Page