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
		<view class="card" v-for="(value,index) in authorList" :key="index">
			<view class="author_avatar_box" v-if="authorList.length!=0">
				<image :src="value.avatar" mode="" class="author_avatar_image"></image>
				<text class="click_me_text_u">{{value.nickName}}</text>
			</view>
			<view class="card_title" @click="linkAuthorDetail(value.authorsId)">
				查看全部>
			</view>
			<view class="card_content" v-if="authorList.length!=0">
				<view class="content_image" @click="linkImgDetail(value.images[0].ID)">
					<image :src="value.images[0].url" mode=""></image>
				</view>
				<view class="content_image" @click="linkImgDetail(value.images[1].ID)">
					<image :src="value.images[1].url" mode=""></image>
				</view><view class="content_image" @click="linkImgDetail(value.images[2].ID)">
					<image :src="value.images[2].url" mode=""></image>
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
	let authorList = ref([]);
	onMounted(() => {
		uni.request({
			url: 'http://127.0.0.1:8888/index/authorsindex', //仅为示例，并非真实接口地址。
			success: (res) => {
				 authorList.value = res.data.data.result
			}
		});
	})
	const linkImgDetail=(imgId)=>{
		uni.navigateTo({
			url: '/pages/index/detail/detail?id='+imgId
		});
	}
	const linkAuthorDetail=(authorId)=>{
		uni.navigateTo({
			url: '/pages/author/detail/detail?id='+authorId
		});
	}
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
