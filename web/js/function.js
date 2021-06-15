axios.defaults.baseURL = 'http://localhost:8083/api/v1';
let app = new Vue({
    el: "#player",
    data: {
        hasLogin:true,
        notLogin:false,
        // 搜索关键字
        query: '周杰伦',
        // 歌曲列表
        musicList: [],
        // 歌曲url
        musicUrl: '',
        // 是否正在播放
        isPlay: false,
        // 歌曲热门评论
        hotComments: [],
        // 歌曲封面地址
        coverUrl: 'https://picture.nj-jay.com/avatar.jpg',
        // 显示视频播放
        showVideo: false,
        // mv地址
        mvUrl: '',
        click: '',
        loginForm:{
            username:'',
            password:''
        },
        messages:[]
    },
    // 方法
    methods: {
        login:function () {
            let param = new URLSearchParams()
            param.append('username', this.loginForm.username)
            param.append('password', this.loginForm.password)
            axios ({
                method:'post',
                url:"/trueLogin",
                data:param
            }).then(function(response){
                console.log(response.data)
                if (response.data.message=="login successfully"){
                    alert("登录成功，刷新页面进入聊天室")
                    //登录成功就跳转
                    this.hasLogin = true;
                    this.notLogin = false;
                    let user_token = response.data.data;
                    let user_name = response.data.username;
                    // 存token到localStorage中
                    localStorage.setItem("currentUser_token", user_token);
                    localStorage.setItem("currentUser_name", user_name);
                } else {
                    this.hasLogin=false;
                    this.notLogin = true;
                    alert("用户名/密码错误");
                }
            })
        },
        registerUser:function () {
            let param = new URLSearchParams()
            param.append('username', this.loginForm.username)
            param.append('password', this.loginForm.password)
            axios ({
                method:'post',
                url:"/adduser",
                data:param
            }).then(function(response){
                console.log(response.data)
                if (response.data.message==="success add user"){
                    alert("注册成功,再刷新页面之后即可进入聊天室")
                    let user_token = response.data.data;
                    let user_name = response.data.username;
                    // 存token到localStorage中
                    localStorage.setItem("currentUser_token", user_token);
                    localStorage.setItem("currentUser_name", user_name);

                } else {
                    this.hasLogin=false;
                    this.notLogin = true;
                    alert("注册失败，该用户已经存在");
                }
            })
        },
        checkLogin:function () {
            let a = localStorage.getItem('currentUser_token')
            if (!a) {
                this.hasLogin = false;
                this.notLogin = true;
            }
        },
        // 搜索歌曲
        searchMusic:function() {
            let that = this;
            axios.get('/search?name=' + this.query)
                .then(function(response) {
                    that.musicList = response.data.data;
                }, function(err){})
        },
        // 播放歌曲
        playMusic:function(musicId) {
            let that = this;
            // 获取歌曲url
            axios.get('/song/url?id=' + musicId)
                .then(function(response) {
                    // 保存歌曲url地址
                    console.log(response.data)
                    that.musicUrl = response.data.data.url;
                })
        },
        // audio的play事件
        play() {
            this.isPlay = true
            // 清空mv的信息
            this.mvUrl = ''
        },
        // audio的pause事件
        pause() {
            this.isPlay = false
        },
        // // 播放mv
        // playMv(vid) {
        //     if (vid) {
        //         this.showVideo = true;
        //         // 获取mv信息
        //         axios.get('/mv/url?id=' + vid).then(response => {
        //             // console.log(response)
        //             // 暂停歌曲播放
        //             this.$refs.audio.pause()
        //             // 获取mv地址
        //             this.mvUrl = response.data.data.url
        //         })
        //     }
        // },
        // 关闭mv界面
        closeMv() {
            this.showVideo = false
            this.$refs.video.pause()
        },
        // // 搜索历史记录中的歌曲
        // historySearch(history) {
        //     this.query = history
        //     this.searchMusic()
        //     this.showHistory = false;
        // },
        clickCount:function(){
            let that = this;
            axios.get('/click')
                .then(function(response) {
                    that.click = response.data.data;
                }, function(err){})
        },
        getCacheMessages:function(){
            let that = this;
            axios.get('/cache')
                .then(function(response) {
                    that.messages = response.data.data;
                })
        }
    },
    mounted(){
        this.clickCount();
        this.searchMusic();
        this.checkLogin();
        this.getCacheMessages();
    }
})
