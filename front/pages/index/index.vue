<template>
	<view class="index">
		<view class="index_banner">

		</view>
		<view class="search">
			<u-search placeholder="请输入创作者口令" v-model="keyword" shape="round" borderColor="#473be5"
				searchIconColor="#473be5"
				:actionStyle="{background:'#473be5',height:'30px',width:'70px',fontSize:'20px',borderRadius:'15px',}"
				style="width: 320px;">
			</u-search>
		</view>
		<view class="author">
			<view class="author_title">
				推荐创作者
			</view>
			<view class="author_content">

				<u-scroll-list :indicator="false">

					<u-avatar v-for="index in 15" :key="index" text="北" fontSize="18" randomBgColor>
					</u-avatar>
				</u-scroll-list>
			</view>
		</view>
		<view class="content">
			<view class="content_img">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</view>
			<view class="content_img">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</view>
			<view class="content_img">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</view>
			<view class="content_img">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</view>
			<view class="content_img">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</view>
			<view class="content_img">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</view>
		</view>
		<u-tabbar :value="tabeIndex" :fixed="true" @change="name => tabeIndex = name" :placeholder="true" :safeAreaInsetBottom="true" zIndex="100">
			<u-tabbar-item text="首页" icon="home" @click="navigate"></u-tabbar-item>
			<u-tabbar-item text="创作者" icon="photo" @click="navigate"></u-tabbar-item>
			<u-tabbar-item text="我的" icon="account" @click="navigate"></u-tabbar-item>
		</u-tabbar>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				keyword: "",
				tips: '',
				value: '',
				title: 'Hello',
				tabeIndex: 0,
				indicator: true,
				list: [{
					thumb: "https://api.lorem.space/image/face?hash=88560"
				}, {
					thumb: "https://api.lorem.space/image/face?hash=88561"
				}, {
					thumb: "https://api.lorem.space/image/face?hash=88562"
				}, {
					thumb: "https://api.lorem.space/image/face?hash=88563"
				}, {
					thumb: "https://api.lorem.space/image/face?hash=88564"
				}],
				baseList: [
					// {
					// 	thumb: "https://api.lorem.space/image/movie?w=80&h=200"
					// },
					// {
					// 	thumb: "https://api.lorem.space/image/movie?w=80&h=200"
					// }, {
					// 	thumb: "https://api.lorem.space/image/movie?w=80&h=200"
					// }, {
					// 	thumb: "https://api.lorem.space/image/movie?w=80&h=200"
					// }, {
					// 	thumb: "https://api.lorem.space/image/movie?w=80&h=200"
					// }, {
					// 	thumb: "https://api.lorem.space/image/movie?w=80&h=200"
					// },
				]
			}
		},
		onLoad() {
			this.tabeIndex = 0
		},
		methods: {
			click1() {
				console.log(123132);
			},
			getCode() {

			},
			navigate() {
				console.log(this.tabeIndex);
				if (this.tabeIndex === 0) {
					return
				}
				const urlArray = ['/pages/index/index', '/pages/author/author', '/pages/me/me']
				uni.navigateTo({
					url: urlArray[this.tabeIndex],
				});
			}
		}
	}
</script>

<style lang="scss">
	.index {
		display: flex;
		flex-direction: column;
		align-items: center;
		height: 100vh;
		width: 100%;
		justify-content: start;
		background-color: green;
	}

	.index_banner {
		height: 500upx;
		background-color: #99a9bf;
		width: 100%;
	}

	.search {
		margin-top: 30upx;
	}

	.author {
		margin-top: 30upx;
		height: 200upx;
		display: flex;
		flex-direction: column;
		justify-content: space-evenly;
		align-items: center;

		.author_title {
			font-size: 40upx;
			color: white;
		}

		.author_content {
			margin-top: 20upx;
			margin-left: 60upx;
			display: flex;
			flex-direction: row;
			justify-content: space-between;
		}
	}

	.u-scroll-list__scroll-view__content {
		.u-avatar {
			margin-left: 10upx;
		}
	}

	.content {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		justify-content: space-around;
		width: 100%;
	}
	.content_img{
		// border: solid;
		// border-color: #99a9bf;
		width: 30%;
	}
</style>
