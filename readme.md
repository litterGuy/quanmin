# ffmpeg

## 合并视频

1. 两个视频分屏,左右播放

```markdown
ffmpeg -i 1.mp4 -i 2.mp4 -filter_complex hstack out.mp4
```

> hstack 横向合并视频
> 
> vstack 纵向合并视频

2. 顺序播放

```markdown
ffmpeg -i 1.mp4 -vcodec copy -acodec copy -vbsf h264_mp4toannexb 1.ts

ffmpeg -i 2.mp4 -vcodec copy -acodec copy -vbsf h264_mp4toannexb 2.ts

ffmpeg -i "concat:1.ts|2.ts" -acodec copy -vcodec copy -absf aac_adtstoasc output.mp4
```

> 先将视频转化为ts文件，然后再进行合并转化为mp4文件
>
> 速度快，文件小

3. 混屏

```markdown
ffmpeg -i 1.mp4 -i 2.mp4 -i 3.mp4 -i 4.mp4 -filter_complex "[0:v]scale=300:200,pad=600:400:0:0[left];[1:v]scale=300:200[right];[left][right]overlay=300:0[up];[2:v]scale=300:200
[down];[up][down]overlay=0:200[downleft];[3:v]scale=300:200[downright];[downleft][downright]overlay=300:200;amix=inputs=4" out.mp4
```

>  命令中首先指定了4路输入；然后按一下步骤进行操作：
> 
>  1.指定第一路流的视频([0:v])作为输入，进行比例变换(scale=300:200)、并填充视频(pad=600:400:0:0),输出为[left]；
> 
>  2.指定第二路流的视频([1:v])为输入进行比例变换(scale=300:200)，输出为[right];
> 
>  3.把[right]叠加到[left]上([left][right]overlay)，并指定位置，坐标为(300:0）,输出为[up]；
> 
>  4.指定第三路流的视频([2:v]) 为输入进行比例变换(scale=300:200),输出为[down]；
> 
>  5. 把[down] 叠加到[up] 上([up][down]overlay)并指定位置，坐标为(0:200),输出为 [downleft]；
> 
>  6.指定第四路流的视频([3:v])为输入进行比例变换(scale=300:200)，输出为[downright];
> 
>  7. 把[downright]叠加到[downleft]上([downleft][downright]overlay) 并指定位置,坐标为(300:200)
> 
>  8. amix=inputs=4，对音频进行混流，这里我们指定混4路音频

4. 网格合并视频

```markdown
ffmpeg -f lavfi -i color=c=black:s=1280x720 -vframes 1 black.png
该命令将创建一张1280*720的图片
```

> 视频个数不一定需要是偶数，如果是奇数，可以用黑色图片来占位

## 音视频

1. 合并音视频

```markdown
ffmpeg -i video.mp4 -i audio.wav -c:v copy -c:a aac -strict experimental output.mp4
```
2. 替换音频

```markdown
ffmpeg -i video.mp4 -i audio.wav -c:v copy -c:a aac -strict experimental
-map 0:v:0 -map 1:a:0 output.mp4
```

3. 获取视频中的音频

```markdown
ffmpeg -i input.mp4 -vn -y -acodec copy output.m4a
```

4. 去掉视频中的音频

```markdown
ffmpeg -i input.mp4 -an output.mp4
```

5. 合并两个音频

```markdown
ffmpeg -i input1.mp3 -i input2.mp3 -filter_complex amerge -ac 2 -c:a libmp3lame -q:a 4 output.mp3
```

## 文字

1. 通过text添加文字

```markdown
ffmpeg -i in.mp4 -vf drawtext=fontcolor=white:fontsize=40:fontfile=simhei.ttf:text='Hello World':x=0:y=100 -vframes 20 -y out.mp4
```

参数说明：[https://www.jianshu.com/p/9d24d81ca199](https://www.jianshu.com/p/9d24d81ca199)