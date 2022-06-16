<template>
	<view class="container">
		<view class="back">
			<navigator hover-class="navigator-hover" open-type="navigateBack">
				<uni-icons type="back" size="30"></uni-icons>
			</navigator>
		</view>
		<view class="feedback">
			<navigator url="/pages/feedback/feedback" hover-class="navigator-hover">
				我要吐槽
			</navigator>
		</view>

		<image src="http://rdgcuypkk.hn-bkt.clouddn.com/1655275252微信图片_20220615144044.jpg" mode=""></image>

		<view class="menu">
			<avatar src="http://rdgcuypkk.hn-bkt.clouddn.com/1655275252微信图片_20220615144044.jpg"></avatar>
			<view class="download">
				<uni-icons type="download" color="white" size="40" @click="download"></uni-icons>
				<view class="">
					下载
				</view>

			</view>

			<view class="download">
				<uni-icons type="star" color="white" size="40"></uni-icons>
				<view class="">
					分享
				</view>

			</view>
		</view>

		<view>
			<!-- 			<uni-popup ref="popup" type="bottom">底部弹出 Popup</uni-popup> -->
			<!-- 普通弹窗 -->
			<uni-popup ref="popup">
				<view class="popup-content" style="background-color: #fff;padding: 30upx auto;">
					<view class="close">
						<uni-icons type="close" size="30" @click="popup.close()"></uni-icons>
					</view>
					<view class="up">
						<view class="up_content">
							<view style="font-size: 38upx;">
								超清无水印原图
							</view>
							<view style="font-size: 28upx;color: gray;margin-top: 10upx;">
								看完广告,即可获取
							</view>
						</view>
						<view class="up_download">
							<button type="default" style="background-color: #4738f3;color: white;" @click="downloadimage">超清下载</button>
						</view>
					</view>
					<view class="down" style="font-size: 28upx;">
						看3次广告，当天可额外获得5次免广告高清下载
					</view>
					
				</view>
			</uni-popup>
		</view>
	</view>
</template>

<script setup>
	import {
		onMounted,
		ref
	} from 'vue'
	import avatar from '../../../components/avatar.vue'
	const popup = ref()
	onMounted(() => {
		console.log(popup);
	})

	const download = () => {
		console.log(popup.value);
		popup.value.open('center')

	}
	
	const downloadimage=() =>{
		console.log(1111);
	}
</script>

<style lang="scss" scoped>
	.back{

		position: fixed;
		width: 60upx;
		height: 60upx;
		top: 40upx;
		left: 30upx;
		z-index: 11;
		border: 1px solid white;
		border-radius: 15px;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.popup-content{
		height: 250upx;
		width: 600upx;
		border-radius: 15px;
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		.close{
			align-self: flex-end;
		}
		.up{
			display: flex;
			flex-direction: row;
			justify-content: space-between;
		}
		.down{
			color: red;
		}
	}
	.container {
		width: 100vw;
		height: 100vh;
		background-color: grey;
		overflow-y: hidden;

		image {
			width: 100%;
			height: 100%;
		}
	}

	.download {
		display: flex;
		flex-direction: column;
		font-size: 32upx;
		align-items: center;
	}

	.menu {
		height: 480upx;
		width: 150upx;
		border: 2px solid white;
		position: fixed;
		border-radius: 15px;
		left: 78%;
		top: 60%;
		z-index: 11;
		box-shadow: x-shadow y-shadow blur spread color inset;
		opacity: 0.85;
		display: flex;
		flex-direction: column;
		justify-content: space-evenly;
		align-items: center;
		color: white;
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
		top: 8%;
		z-index: 10;
	}
</style>
