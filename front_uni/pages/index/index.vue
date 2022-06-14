<template>
	<view class="container">
		<view class="feedback" @click="">
			<navigator url="/pages/feedback/feedback" hover-class="navigator-hover">
				我要吐槽
			</navigator>

		</view>
		<view class="banner">

		</view>
		<view class="search">
			<view class="search_bar">
				<uni-icons type="search" size="30" color="rgb(47, 38, 155)"></uni-icons>
				<input type="text">
				<button class="search_btn">搜索</button>
			</view>
		</view>
		<view class="author">
			<view class="author_title">
				推荐创作者
			</view>

			<scroll-view scroll-x="true" scroll-y="false" @scroll="scroll" class="author_scroll">
				<template v-if="authorList.length!=0">
					<view class="author_avatar_box" v-for="(item,index) in authorList" :key="index">
						<image :src="item.avatar" class="author_avatar_image"></image>
					</view>
				</template>
			</scroll-view>
		</view>
		<view class="content">
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
			<view class="content_image">
				<image src="../../static/c1.png" mode=""></image>
			</view>
		</view>
	</view>
</template>

<script setup>
	import {
		reactive,
		ref,
		onMounted,
		nextTick,
		onUnmounted
	} from 'vue'
	let href = 'https://uniapp.dcloud.io/component/README?id=uniui'
	let old = {
		screenTop: 0
	}
	let authorList = ref([])

	const scroll = e => {
		old.scrollTop = e.detail.scrollTop
	}

	onMounted(() => {
		uni.request({
			url: 'http://127.0.0.1:8888/index/index', //仅为示例，并非真实接口地址。
			success: (res) => {
				authorList.value = res.data.data.authors
				console.log(authorList);
			}
		});
	})
	// nextTick(() => {
	// 	console.log('nextTick');
	// 	console.log(authorList);
	// })
</script>

<style scoped lang="scss">
	.container {
		display: flex;
		flex-direction: column;
		width: 100vw;

		background-color: coral;
	}

	.feedback {
		height: 140upx;
		width: 40upx;
		font-size: 25upx;
		font-weight: bold;
		color: white;
		display: inline-block;
		background-color: #4738f3;
		border-top-right-radius: 15upx;
		border-bottom-right-radius: 15upx;
		position: fixed;
		top: 40upx;
	}

	.banner {
		height: 400upx;
		background-color: bisque;
	}

	.search {
		height: 200upx;
		background-color: aliceblue;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.search_bar {
		border: 2upx solid purple;
		border-radius: 25upx;
		height: 100upx;
		width: 90%;
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
	}

	.search_btn {
		width: 200upx;
		color: white;
		background-color: rgb(47, 38, 155);

	}

	.author {
		height: 200upx;
		color: white;
		background-color: chartreuse;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-around;
		/* padding: 20upx; */
	}

	.author_scroll {
		white-space: nowrap;
		width: 100%;
	}

	.author_avatar_box {
		display: inline-block;
		width: 80upx;
		height: 80upx;
		margin-left: 30upx;
	}

	.author_avatar_image {
		width: 80upx;
		height: 80upx;
		border-radius: 50%;
		border: solid 5upx white;
	}

	.content {
		padding: 20upx;
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		gap: 10upx;
	}

	.content_image image {
		height: 200upx;
		width: 100%;

	}
</style>
