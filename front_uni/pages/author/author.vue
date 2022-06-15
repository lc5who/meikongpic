<template>
	<view class="container">
		<view class="banner">
			
		</view>
		
		<view class="card">
			<view class="author_avatar_box">
				<image src="../../static/c2.png" mode="" class="author_avatar_image"></image>
				<text>星月</text>
			</view>
			<view class="card_title">
				查看全部>
			</view>
			<view class="card_content">
				<view class="content_image">
					<image src="../../static/c1.png" mode=""></image>
				</view>
				
				<view class="content_image">
					<image src="../../static/c1.png" mode=""></image>
				</view><view class="content_image">
					<image src="../../static/c1.png" mode=""></image>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>

</script>

<style lang="scss">
	.container {
		width: 100vw;
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		align-items: center;
		background-color: grey;
	}
	.banner{
		height: 200upx;
		width: 98%;
		border: 2upx solid white;
		border-radius: 25upx;
		
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
		border: 4upx solid white;
		border-radius: 25px;
		image{
			width: 100%;
			height: 350upx;
		}
	}
	.author_avatar_box {
		
		width: 200upx;
		height: 80upx;
		position: relative;
		top: -40upx;
		left: 60upx;
	}
	
	.author_avatar_image {
		display: inline-block;
		width: 80upx;
		height: 80upx;
		border-radius: 50%;
		border: solid 5upx white;
	}
</style>
