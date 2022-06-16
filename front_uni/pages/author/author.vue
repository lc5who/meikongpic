<template>
	<view class="container">
		<view class="author_avatar_box_u">
			<view>
			<image src="../../static/logo2.jpg" mode="" class="author_avatar_image_u"></image>
			</view>
			<view>
				<text @click="apitest()" class="click_me_text_u">&nbsp;我的</text>
			</view>
		</view>
		<view class="banner">
			<image src="../../static/join.jpg" class="bannerImage"></image>
		</view>
		<view class="card" v-for="(item,index) in cardArr" :value="index">
			<view class="author_avatar_box" v-if="authorList.length!=0">
				<image :src="authorList[0].avatar" mode="" class="author_avatar_image"></image>
				<text class="click_me_text_u">{{authorList[0].nickname}}</text>
			</view>
			<view class="card_title">
				查看全部>
			</view>
			<view class="card_content" v-if="imageList.length!=0">
				<view class="content_image">
					<image :src="imageList[0].url" mode=""></image>
				</view>
				<view class="content_image">
					<image :src="imageList[0].url" mode=""></image>
				</view><view class="content_image">
					<image :src="imageList[0].url" mode=""></image>
				</view>
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
	const cardArr=[1,2,3,4];
	let imageList = ref([]);
	let authorList = ref([]);
	onMounted(() => {
		uni.request({
			url: 'http://127.0.0.1:8888/index/index', //仅为示例，并非真实接口地址。
			success: (res) => {
	
				 authorList.value = res.data.data.authors
				 imageList.value = res.data.data.imagesList
				// imageList.value = res.data.data.imgs
				console.log(authorList)
			}
		});
	})
</script>

<style lang="scss">
	.container {
		width: 100vw;
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		align-items: center;
		background-color: rgb(20, 20, 58);
	}
	.banner{
		height: 200upx;
		width: 98%;
		// border: 2upx solid white;
		// border-radius: 25upx;
		
	}
	.bannerImage{
		height: 200upx;
		width: 98%;
		// border: 2upx solid white;
		border-radius: 40upx;
		
	}
	.card{
		margin-top: 60upx;
		height: 500upx;
		width: 98%;
		border: 2upx solid white;
		border-radius: 25upx;
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		// margin: 0 10upx 0 10upx;
	}
	.card_title{
		color: white;
		display: flex;
		justify-content: flex-end;
		padding-right: 30upx;
	}
	.card_content{
		display: flex;
		width: 100%;
		flex-direction: row;
		justify-content: space-around;
	}
	.content_image{
		width: 30%;
		image{
			width: 100%;
			height: 350upx;
			border: 4upx solid white;
			border-radius: 25px;
		}
	}
	.author_avatar_box {
		
		width: 200upx;
		height: 80upx;
		position: relative;
		top: 20upx;
		left: 10upx;
		display: flex;
		align-items: center;
		justify-content: flex-start;
	}
	
	.author_avatar_image {
		display: inline-block;
		width: 80upx;
		height: 80upx;
		border-radius: 50%;
		border: solid 5upx white;
	}
	.author_avatar_box_u {
			width: 160upx;
			height: 70upx;
			display: flex;
			align-items: center;
			justify-content: flex-start;
			margin-right: 580upx;
			margin-bottom: 20upx;
		}
		.author_avatar_image_u {
			width: 50upx;
			height: 50upx;
			border-radius: 50%;
			background-color: white;
		}
		.click_me_text_u{
			color: white;
		}
</style>
