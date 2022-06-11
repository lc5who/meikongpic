<template>
	<view class="container">
		<view class="top_bar">
			<uni-icons type="back" size="30" @click="goback"></uni-icons>
			我要吐槽
		</view>

		<view class="tip">
			<uni-icons type="help" size="30"></uni-icons>
			<view class="" style="margin-left: 20upx;">
				常见问题
			</view>
		</view>

		<view class="content">
			<textarea name="content" id="content" cols="30" rows="10" placeholder="请填写你想吐槽的问题,以便我们提供更好的帮助"
				style="width: 100%;height: 90%;">

			</textarea>
		</view>

		<view class="upload_image">
			<view class="label">
				相关截图(选填)(0/4)
			</view>
			<view class="uploder">
				<uni-file-picker limit="4" ></uni-file-picker>
			</view>

		</view>
		<view class="phone">
			<view class="label">联系方式(选填)</view>
			<input class="uni-input phone_input"  placeholder="手机号或邮箱" />
		</view>
		<view class="submit">
				<button size="default" class="btn" @click="submit()">确定提交</button>
			</view>
	</view>
</template>

<script setup>
	let goback= ()=>{
		uni.navigateBack()
	}
	
	let submit = ()=>{
		
	}
</script>

<style lang="scss">
	.container {
		display: flex;
		flex-direction: column;
		width: 100vw;
		height: 100vh;
		justify-content: flex-start;
		background-color: #cacaca;
		align-items: center;
	}

	.top_bar {
		width: 100%;
		background-color: white;
		height: 60upx;
		display: flex;
		justify-self: flex-start;
		align-items: center;
		padding-left: 20upx;
		font-weight: bold;
		font-size: 30upx;
	}

	.tip {
		display: flex;
		width: calc(96% - 20upx);
		height: 60upx;
		background-color: white;
		align-items: center;
		border-radius: 15upx;
		padding-left: 20upx;
		margin-top: 20upx;
	}

	.content {
		margin-top: 20upx;
		height: 400upx;
		width: calc(96% - 20upx);
		background-color: white;
		border-radius: 15upx;
		padding-top: 20upx;
		padding-left: 20upx;
	}

	.upload_image {
		width: 96%;
		display: flex;
		justify-content: flex-start;
		margin-top: 20upx;
		flex-direction: column;
		
		.uploder {
			margin-top: 4upx;
			width: 60%;
			border-color: black;
			color: black;
			
		}
	}

	.label {
		font-size: 26upx;
		font-weight: bold;
	}

	.phone {
		margin-top: 20upx;
		width: 96%;
		height: 60upx;

	}
	.phone_input{
		margin-top: 4upx;
		background-color: white;
		height: 60upx;
		border-radius: 15upx;
		padding-left: 10upx;
	}
	
	.submit{
		margin-top: 80upx;
		width: 80%;
		
		.btn{
			background-color: #7641e9;
			color: white;
			border-radius: 30upx;
		}
	}

</style>
