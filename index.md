# Bmoji - BiliBili表情包浏览器

一款B站表情包浏览器，让你可以打包下载BiliBili表情~

# 使用说明

访问网页[bmoji.sakurasen.cn](https://bmoji.sakurasen.cn/)

# 自部署说明

`/back`目录中为后端（使用Golang编写）

`/vue`目录中为前端，使用Vue3编写

如需自部署，请编译完毕后运行一次生成config.json

然后再填写config.json中的信息，

```
{
    "app": {
        "cors": ["http://localhost:5173"],
        "savetype": "local"
    },
    "bilibili": {
        "sessdata": ""
    }
}
```

cors中填写网站域名，sessdata填写B站cookie中sessdata值,savetype填写local(还没开发更多存储)，都是必填的

然后前往`vue/src`中 `App.vue`,`pages/p_info.vue`修改baseurl的变量为你自己的变量

这是我第一次写开源项目，欢迎指错。