axios.defaults.baseURL = 'http://localhost:8081/api/v1';

let app = new Vue({
    el: "#player",
    data: {
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
        click: ''
    },
    // 方法
    methods: {
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
        }
    },
    mounted(){
        this.clickCount();
        this.searchMusic();
    }
})