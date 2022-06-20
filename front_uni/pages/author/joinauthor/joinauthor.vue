<template>
	<view class="container" style="position：relative">
		<view class="top_bar">
			<uni-icons type="back" size="30" @click="goback"></uni-icons>
			成为创作者
		</view>
		<image src="../../../static/joinAuthor.jpg" style="height: 1200px;width: 100%;"></image>
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
}
</style>
