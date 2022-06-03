# 基于gin搭建的web项目骨架



<!-- GETTING STARTED -->
## 🏗 快速开始

### 环境要求

* Go (1.17+)
* MySQL (5.7+)
* Redis


### 运行说明


克隆代码库

   ```sh
   git clone https://github.com/xiaoriri-team/gin-skeleton.git
   ```

#### 后端

1. 导入项目根目录下的 `init.sql` 文件至MySQL数据库
2. 拷贝项目根目录下 `config.yaml.sample` 文件至 `config.yaml`，按照注释完成配置编辑
3. 编译后端

    ```sh
    go mod download
    go build -o gin-skeleton .
    ```

4. 启动后端

    ```sh
    chmod +x gin-skeleton
    ./gin-skeleton
    ```

### 其他说明

建议后端服务使用 `supervisor` 守护进程，并通过 `nginx` 反向代理后，提供API给前端服务调用。


代码结构比较简单，很方便扩展

## 👯‍♀️ 贡献

喜欢的朋友欢迎给个Star、贡献PR。

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/xiaoriri-team/gin-skeleton?style=flat
[contributors-url]: https://github.com/xiaoriri-team/gin-skeleton/graphs/contributors
[goreport-shield]: https://goreportcard.com/badge/github.com/xiaoriri-team/gin-skeleton
[goreport-url]: https://goreportcard.com/report/github.com/xiaoriri-team/gin-skeleton
[forks-shield]: https://img.shields.io/github/forks/xiaoriri-team/gin-skeleton?style=flat
[forks-url]: https://github.com/xiaoriri-team/gin-skeleton/network/members
[stars-shield]: https://img.shields.io/github/stars/xiaoriri-team/gin-skeleton.svg?style=flat
[stars-url]: https://github.com/xiaoriri-team/gin-skeleton/stargazers
[issues-shield]: https://img.shields.io/github/issues/xiaoriri-team/gin-skeleton.svg?style=flat
[issues-url]: https://github.com/xiaoriri-team/gin-skeleton/issues
[license-shield]: https://img.shields.io/github/license/xiaoriri-team/gin-skeleton.svg?style=flat
[license-url]: https://github.com/xiaoriri-team/gin-skeleton/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat&logo=linkedin&colorB=555
