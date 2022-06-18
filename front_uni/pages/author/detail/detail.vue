<template>
	<view class="container">
		<view class="back">
			<navigator hover-class="navigator-hover" open-type="navigateBack">
				<uni-icons type="back" size="30"></uni-icons>
			</navigator>
		</view>
		<view class="author_info">
			<view class="author_up">
				<avatar src="http://rdgcuypkk.hn-bkt.clouddn.com/1655185007logo.png" size="70"></avatar>

				<view class="author_name">
					{{authorData.nickname}}
				</view>


			</view>
			<view class="author_down">
				{{authorData.info}}
			</view>
		</view>
		<view class="image_list">
			<view class="header">
				<view class="zuopin_btn">
					作品(9)
				</view>

			</view>
			<hr style="border:1upx solid grey;margin-top: 10upx;" />
			<!-- <scroll-view scroll-y="true" @scroll="scroll" class="content"> -->
			<view class="content">
				<navigator :url="'/pages/index/detail/detail?id='+item.ID" v-for="(item,index) in imageData" :key="index">
					<view class="content_image">
						<image :src="item.url" mode=""></image>
					</view>
				</navigator>
			</view>
				
			<!-- </scroll-view> -->

		

		</view>
		<view class="footer">
			<view class="line">
				<hr style="border: 1upx solid darkgrey;" />
			</view>
			我们是有底线的
			<view class="line">
				<hr style="border: 1upx solid darkgrey;" />
			</view>
		</view>
	</view>
</template>

<script setup>
	import {
		onMounted,
		ref,
		reactive
	} from 'vue'
	import avatar from '../../../components/avatar.vue'
	let imageData = ref({})
	let authorData = ref({})
	let old = {
		screenTop: 0
	}
	const scroll = e => {
		old.scrollTop = e.detail.scrollTop
	}
	onMounted(() => {
		let routes = getCurrentPages(); // 获取当前打开过的页面路由数组
		let curRoute = routes[routes.length - 1].route //获取当前页面路由
		let curParam = routes[routes.length - 1].$page.options; //获取路由参数
		uni.request({
			url: 'http://127.0.0.1:8888/index/authorDetail?id=' + curParam.id, //仅为示例，并非真实接口地址。
			success: (res) => {
				imageData.value = res.data.data.image
				authorData.value = res.data.data.author
				// console.log(imageData);
				console.log(authorData);
			}
		});

	})
</script>

<style lang="scss">
	.back {
		position: fixed;
		width: 60upx;
		height: 60upx;
		top: 30upx;
		left: 20upx;
		z-index: 11;
		border: 1px solid darkgray;
		border-radius: 30upx;
		display: flex;
		justify-content: center;
		align-items: center;

	}

	.container {
		width: 100vw;
		display: flex;
		flex-direction: column;
		background-color: #000;
		padding-top: 120upx;
		height: 100vh;
		overflow-y: auto;
	}

	.author_info {
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		color: white;
		padding-left: 30upx;
		padding-right: 30upx;
	}

	.author_up {
		display: flex;
		flex-direction: row;

		align-items: center;

		.author_name {
			margin-left: 20upx;
			font-size: 28upx;
			font-weight: bold;
		}
	}

	.author_down {
		margin-top: 22upx;
		display: block;
		font-size: 22upx;
	}

	.footer {
		background-color: #000;
		margin-top: 20upx;
		color: darkgrey;
		display: flex;
		justify-content: space-around;
		align-items: center;
		font-size: 22upx;

		.line {
			width: 26%;
		}
	}

	.content {
		padding-top: 20upx;
		padding-left: 2%;
		padding-right: 2%;
		width: 100%;
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		gap: 20upx;
		background-color: #000;
	}

	.content_image {
		height: 300upx;

	}

	.content_image image {
		height: 100%;
		width: 100%;
		border: 1px solid #fff;
		border-radius: 10upx;

	}

	.header {
		margin-top: 30upx;
		padding-left: 40upx;

		.zuopin_btn {
			height: 50upx;
			width: 150upx;
			background-color: #4738f3;
			border-radius: 20upx;
			text-align: center;
			line-height: 50upx;
		}
	}
</style>
