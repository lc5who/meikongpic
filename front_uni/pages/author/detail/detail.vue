<template>
	<view class="container">
		<view class="author_info">
			<view class="author_up">
				<avatar src="http://rdgcuypkk.hn-bkt.clouddn.com/1655185007logo.png" size="60"></avatar>
				<view class="author_name">
					你爹
				</view>
			</view>
			<view class="author_down">
				我是你爹嘻嘻嘻嘻我是你爹嘻嘻嘻嘻我是你爹嘻嘻嘻嘻我是你爹嘻嘻嘻嘻我是你爹嘻嘻嘻嘻我是你爹嘻嘻嘻嘻我是你爹嘻嘻嘻嘻
			</view>
		</view>
		<view class="image_list">
			<view class="header">
				<uni-tag :circle="true" :inverted="inverted" text="作品" type="primary" @click="setInverted" />
			</view>
		</view>
	</view>
</template>

<script setup>
	import {
		onMounted,
		ref
	} from 'vue'
	import avatar from '../../../components/avatar.vue'
	onMounted(()=>{
		let routes = getCurrentPages(); // 获取当前打开过的页面路由数组
		let curRoute = routes[routes.length - 1].route //获取当前页面路由
		let curParam = routes[routes.length - 1].$page.options; //获取路由参数
		console.log(curParam.id);
		let imageData=ref({})
		let authorData=ref({})
		uni.request({
			url: 'http://127.0.0.1:8888/index/authorDetail?id='+curParam.id, //仅为示例，并非真实接口地址。
			success: (res) => {
				console.log(res.data.data)
				imageData.value =res.data.data.image
				authorData.value  =res.data.data.author
				console.log(imageData.value);
				console.log(authorData.value);
			}
		});
		
	})
</script>

<style lang="scss">
	.container {
		width: 100vw;
		display: flex;
		flex-direction: column;
		background-color: #000;
		padding-top: 60upx;
	}
	.author_info{
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		color: white;
		padding-left: 30upx;
		padding-right: 30upx;
	}
	
	.author_up{
		display: flex;
		flex-direction: row;
		
		align-items: center;
		
		.author_name{
			margin-left: 20upx;
			font-size: 28upx;
			font-weight: bold;
		}
	}
	.author_down{
		margin-top: 22upx;
		display: block;
		font-size: 22upx;
	}
	
</style>
