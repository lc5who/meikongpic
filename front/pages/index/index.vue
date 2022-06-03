<template>
	<view class="u-page" style="height: 100vh;width: 100%;background-color: green; padding-left: 5%;padding-right: 5%;">
		<u-search placeholder="请输入创作者口令" v-model="keyword" shape="round" borderColor="#473be5" searchIconColor="#473be5" :actionStyle="{background:'#473be5',height:'30px',width:'70px',fontSize:'20px',borderRadius:'15px',}" style="width: 320px;">
		</u-search>
		<view class="">
			推荐创作者
		</view>
		<u-scroll-list>
			<view v-for="(item, index) in list" :key="index">
				<u-avatar :src="item.thumb"></u-avatar>
			</view>
		</u-scroll-list>
		
		<u-row justify="space-between" gutter="10">
			<u-col span="4">
				<view class="" style="border-radius: 1.5rem;border-color: white;border-width: 2upx;">
					<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
						height="200px" @click="click"></u--image>
				</view>
			</u-col>
			<u-col span="4">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</u-col>
			<u-col span="4">
				<u--image :showLoading="true" src="https://api.lorem.space/image/movie?w=80&h=200" width="100%"
					height="200px" @click="click" radius="15upx"></u--image>
			</u-col>
		</u-row>
		<!-- 		<u-grid :border="false" col="3">
			<u-grid-item v-for="(baseListItem,baseListIndex) in baseList" :key="baseListIndex">
				<u--image :showLoading="true" :src="baseListItem.thumb" width="80px" height="200px" @click="click"></u--image>
			</u-grid-item>
		</u-grid> -->
		<u-tabbar :value="tabeIndex" :fixed="true" :placeholder="true"
			:safeAreaInsetBottom="true" zIndex="100">
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
				tabeIndex:0,
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
			this.tabeIndex=0
		},
		methods: {
			click1() {
				console.log(123132);
			},
			getCode() {

			},
			navigate(){
				console.log(this.tabeIndex);
				if(this.tabeIndex===0){
					return
				}
				const urlArray=['/pages/index/index','/pages/author/author','/pages/me/me']
				uni.navigateTo({
					url: urlArray[this.tabeIndex],
				});
			}
		}
	}
</script>

<style lang="scss">
	.content {
		display: flex;
		flex-direction: column;

	}

	.wrap {
		padding: 12px;
	}

	.demo-layout {
		height: 25px;
		border-radius: 4px;
	}

	.bg-purple {
		background: #CED7E1;
	}

	.bg-purple-light {
		background: #e5e9f2;
	}

	.bg-purple-dark {
		background: #99a9bf;
	}
</style>
