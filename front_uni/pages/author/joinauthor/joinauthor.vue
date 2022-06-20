<template>
	<view class="container" style="position：relative">
		<view class="top_bar">
			<view>
			<uni-icons type="back" size="30" @click="goback"></uni-icons>
			</view>
			<view>
			<text>成为创作者</text>
			</view>
		</view>
		<image src="../../../static/joinAuthor.jpg"></image>
	</view>
</template>

<script setup>
	let goback= ()=>{
		uni.navigateBack()
	}
</script>

<style lang="scss">
.top_bar{
	position: absolute;
	z-index: 999;
	background: transparent;
	display: flex;
	align-items: center;
	justify-content: flex-start;
	text{
		color: rgba(255,255,255,0.7);
		font-weight: 900;
	}
}
image{
	height: 1200px;
	width: 100%;
}
</style>
