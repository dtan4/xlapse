load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
        version = "v0.0.0-20190523083050-ea95bdfd59fc",
    )
    go_repository(
        name = "com_github_ajg_form",
        importpath = "github.com/ajg/form",
        sum = "h1:t9c7v8JUKu/XxOGBU0yjNpaMloxGEJhUkqFRq0ibGeU=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_andreasbriese_bbloom",
        importpath = "github.com/AndreasBriese/bbloom",
        sum = "h1:HD8gA2tkByhMAwYaFAX9w2l7vxvBQ5NMoxDrkhqhtn4=",
        version = "v0.0.0-20190306092124-e2d15f34fcf9",
    )
    go_repository(
        name = "com_github_andybalholm_brotli",
        importpath = "github.com/andybalholm/brotli",
        sum = "h1:V7DdXeJtZscaqfNuAdSRuRFzuiKlHSC/Zh3zl9qY3JY=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_antihax_optional",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_armon_consul_api",
        importpath = "github.com/armon/consul-api",
        sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
        version = "v0.0.0-20180202201655-eb2c6b5be1b6",
    )
    go_repository(
        name = "com_github_aws_aws_lambda_go",
        importpath = "github.com/aws/aws-lambda-go",
        sum = "h1:qelHgOUidrQmrfFTLiC7u6wWuuwBJ9yKcjVRkIy7834=",
        version = "v1.30.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:vPlF4cUsdN5ETfvb7ewZFbFZyB6Rsfndt3kS2XqLXKo=",
        version = "v1.42.44",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:fqlCk6Iy3bnCumtrLz9r3mJ/2gUT0pJ0wLFVIdWh+JA=",
        version = "v1.16.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_aws_protocol_eventstream",
        importpath = "github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream",
        sum = "h1:SdK4Ppk5IzLs64ZMvr6MrSficMtjY2oS0WOORXTlxwU=",
        version = "v1.4.1",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_config",
        importpath = "github.com/aws/aws-sdk-go-v2/config",
        sum = "h1:5AlQD0jhVXlGzwo+VORKiUuogkG7pQcLJNzIzK7eodw=",
        version = "v1.15.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:RQQ5fzclAKJyY5TvF+fkjJEwzK4hnxQCLOu5JXzDmQo=",
        version = "v1.11.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
        sum = "h1:LWPg5zjHV9oz/myQr4wMs0gi4CjnDN/ILmyZUFYXZsU=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
        sum = "h1:onz/VaaxZ7Z4V+WIN9Txly9XLTmoOh1oJ8XcAC3pako=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
        sum = "h1:9stUQR/u2KXU6HkFJYlqnZEjBnbgrVbG6I5HN09xZh0=",
        version = "v2.4.3",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_ini",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/ini",
        sum = "h1:by9P+oy3P/CwggN4ClnW2D4oL91QV7pBzBICi1chZvQ=",
        version = "v1.3.10",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_v4a",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/v4a",
        sum = "h1:cq+47u1zpHyH+PSkbBx1N9whx4TiM9m9ibimOPaNlBg=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_accept_encoding",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
        sum = "h1:T4pFel53bkHjL2mMo+4DKE6r6AuoZnM0fg7k1/ratr4=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_checksum",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/checksum",
        sum = "h1:I0dcwWitE752hVSMrsLCxqNQ+UdEp3nACx2bYNMQq+k=",
        version = "v1.1.3",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:Gh1Gpyh01Yvn7ilO/b/hr01WgNpaszfbKMUgqM186xQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_s3shared",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
        sum = "h1:BKjwCJPnANbkwQ8vzSbaZDKawwagDubrH/z/c0X+kbQ=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_lambda",
        importpath = "github.com/aws/aws-sdk-go-v2/service/lambda",
        sum = "h1:y4iAiwisIY1FKyaoYU6hHVjQ2N0/4aUWow0021/acXU=",
        version = "v1.22.0",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_route53",
        importpath = "github.com/aws/aws-sdk-go-v2/service/route53",
        sum = "h1:OsggywXCk9iFKdu2Aopg3e1oJITIuyW36hA/B0rqupE=",
        version = "v1.6.2",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_s3",
        importpath = "github.com/aws/aws-sdk-go-v2/service/s3",
        sum = "h1:A3PuAUlh1u47WHcM68CDaG9ZWjK7ewePjDp+0dY9yv4=",
        version = "v1.26.5",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sso",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
        sum = "h1:frW4ikGcxfAEDfmQqWgMLp+F1n4nRo9sF39OcIb5BkQ=",
        version = "v1.11.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sts",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
        sum = "h1:cJGRyzCSVwZC7zZZ1xbx9m32UnrKydRYhOvcD1NYP9Q=",
        version = "v1.16.3",
    )
    go_repository(
        name = "com_github_aws_aws_xray_sdk_go",
        importpath = "github.com/aws/aws-xray-sdk-go",
        sum = "h1:mATj8779Kj8Ae8oyXZ3S4GeK9BDGHqlLAKGCUiE31o4=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:eG/N+CcUMAvsdffgMvjMKwfyDzIkjM6pfxMJ8Mzc6mE=",
        version = "v1.11.2",
    )
    go_repository(
        name = "com_github_aymerick_raymond",
        importpath = "github.com/aymerick/raymond",
        sum = "h1:Ppm0npCCsmuR9oQaBtRuZcmILVE74aXE+AmrJj8L2ns=",
        version = "v2.0.3-0.20180322193309-b565731e1464+incompatible",
    )
    go_repository(
        name = "com_github_benbjohnson_clock",
        importpath = "github.com/benbjohnson/clock",
        sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_burntsushi_xgb",
        importpath = "github.com/BurntSushi/xgb",
        sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version = "v0.0.0-20160522181843-27f122750802",
    )

    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash",
        importpath = "github.com/cespare/xxhash",
        sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
        version = "v2.1.1",
    )

    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cloudykit_fastprinter",
        importpath = "github.com/CloudyKit/fastprinter",
        sum = "h1:sR+/8Yb4slttB4vD+b9btVEnWgL3Q00OBTzVT8B9C0c=",
        version = "v0.0.0-20200109182630-33d98a066a53",
    )
    go_repository(
        name = "com_github_cloudykit_jet",
        importpath = "github.com/CloudyKit/jet",
        sum = "h1:rZgFj+Gtf3NMi/U5FvCvhzaxzW/TaPYgUYx3bAPz9DE=",
        version = "v2.1.3-0.20180809161101-62edd43e4f88+incompatible",
    )
    go_repository(
        name = "com_github_cloudykit_jet_v3",
        importpath = "github.com/CloudyKit/jet/v3",
        sum = "h1:1PwO5w5VCtlUUl+KTOBsTGZlhjWkcybsGaAau52tOy8=",
        version = "v3.0.0",
    )

    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:hzAQntlaYRkVSFEfj9OTWlVV1H155FMD8BTKktLv0QI=",
        version = "v0.0.0-20210930031921-04548b0d99d4",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:zH8ljVhhq7yC0MIeUL/IviMtY8hx2mK8cN9wEYb8ggw=",
        version = "v0.0.0-20211011173535-cb28da3451f1",
    )

    go_repository(
        name = "com_github_codegangsta_inject",
        importpath = "github.com/codegangsta/inject",
        sum = "h1:sDMmm+q/3+BukdIpxwO365v/Rbspp2Nt5XntgQRXq8Q=",
        version = "v0.0.0-20150114235600-33e0aa1cb7c0",
    )
    go_repository(
        name = "com_github_coreos_etcd",
        importpath = "github.com/coreos/etcd",
        sum = "h1:jFneRYjIvLMLhDLCzuTuU4rSJUjRplcJQ7pD7MnhC04=",
        version = "v3.3.10+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_etcd",
        importpath = "github.com/coreos/go-etcd",
        sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:3Jm3tLmsgAYcjC+4Up7hJrFBPr+n7rAqYeSw/SZazuY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man",
        importpath = "github.com/cpuguy83/go-md2man",
        sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
        version = "v1.0.10",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_data_dog_go_sqlmock",
        importpath = "github.com/DATA-DOG/go-sqlmock",
        sum = "h1:ThlnYciV1iM/V0OSF/dtkqWb6xo5qITT1TJBG1MRDJM=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_dgraph_io_badger",
        importpath = "github.com/dgraph-io/badger",
        sum = "h1:DshxFxZWXUcO0xX476VJC07Xsr6ZCBVRHKZ93Oh7Evo=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        importpath = "github.com/dgrijalva/jwt-go",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_dgryski_go_farm",
        importpath = "github.com/dgryski/go-farm",
        sum = "h1:tdlZCpZ/P9DhczCTSixgIKmwPv6+wP5DGjqLYw5SUiA=",
        version = "v0.0.0-20190423205320-6a90982ecee2",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_eknkc_amber",
        importpath = "github.com/eknkc/amber",
        sum = "h1:clC1lXBpe2kTj2VHdaIu9ajZQe4kcEY9j0NsnDDBZ3o=",
        version = "v0.0.0-20171010120322-cdade1c07385",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:fP+fF0up6oPY49OrjPrhIJ8yQfdIM85NXMLkMg1EXVs=",
        version = "v0.9.10-0.20210907150352-cf90f659a021",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_etcd_io_bbolt",
        importpath = "github.com/etcd-io/bbolt",
        sum = "h1:gSJmxrs37LgTqR/oyJBWok6k6SvXEUerFTbltIhXkBM=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_github_fasthttp_contrib_websocket",
        importpath = "github.com/fasthttp-contrib/websocket",
        sum = "h1:DddqAaWDpywytcG8w/qoQ5sAN8X12d3Z3koB0C3Rxsc=",
        version = "v0.0.0-20160511215533-1f3b11f56072",
    )
    go_repository(
        name = "com_github_fatih_structs",
        importpath = "github.com/fatih/structs",
        sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_flosch_pongo2",
        importpath = "github.com/flosch/pongo2",
        sum = "h1:GY1+t5Dr9OKADM64SYnQjw/w99HMYvQ0A8/JoUkxVmc=",
        version = "v0.0.0-20190707114632-bbf5a6c351f4",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
        version = "v1.4.7",
    )
    go_repository(
        name = "com_github_gavv_httpexpect",
        importpath = "github.com/gavv/httpexpect",
        sum = "h1:1X9kcRshkSKEjNJJxX9Y9mQ5BRfbxU5kORdjhlA1yX8=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_getsentry_sentry_go",
        importpath = "github.com/getsentry/sentry-go",
        sum = "h1:20dgTiUSfxRB/EhMPtxcL9ZEbM1ZdR+W/7f7NWD+xWo=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_gin_contrib_sse",
        importpath = "github.com/gin-contrib/sse",
        sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gin_gonic_gin",
        importpath = "github.com/gin-gonic/gin",
        sum = "h1:3DoBmSbJbZAWqXJC3SLjAPfutPJJRN1U5pALB7EeTTs=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_github_go_check_check",
        importpath = "github.com/go-check/check",
        sum = "h1:0gkP6mzaMqkmpcJYCFOLkIBwI7xFExG03bbkOkCvUPI=",
        version = "v0.0.0-20180628173108-788fd7840127",
    )
    go_repository(
        name = "com_github_go_errors_errors",
        importpath = "github.com/go-errors/errors",
        sum = "h1:LUHzmkK3GUKUrL/1gfBUxAHzcev3apQlezX/+O7ma6w=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_go_gl_glfw_v3_3_glfw",
        importpath = "github.com/go-gl/glfw/v3.3/glfw",
        sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
        version = "v0.0.0-20200222043503-6f7a984d4dc4",
    )
    go_repository(
        name = "com_github_go_kit_kit",
        importpath = "github.com/go-kit/kit",
        sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:MP4Eh7ZCb31lleYCFuwm0oe4/YGak+5l1vA2NOE80nA=",
        version = "v0.4.0",
    )

    go_repository(
        name = "com_github_go_martini_martini",
        importpath = "github.com/go-martini/martini",
        sum = "h1:xveKWz2iaueeTaUgdetzel+U7exyigDYBryyVfV/rZk=",
        version = "v0.0.0-20170121215854-22fa46961aab",
    )
    go_repository(
        name = "com_github_go_playground_assert_v2",
        importpath = "github.com/go-playground/assert/v2",
        sum = "h1:MsBgLAaY856+nPRTKrp3/OZK38U/wa0CcBYNjji3q3A=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_go_playground_locales",
        importpath = "github.com/go-playground/locales",
        sum = "h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_go_playground_universal_translator",
        importpath = "github.com/go-playground/universal-translator",
        sum = "h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=",
        version = "v0.17.0",
    )
    go_repository(
        name = "com_github_go_playground_validator_v10",
        importpath = "github.com/go-playground/validator/v10",
        sum = "h1:pH2c5ADXtd66mxoE0Zm9SUhxE20r7aM3F26W0hOn+GE=",
        version = "v10.4.1",
    )

    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )

    go_repository(
        name = "com_github_gobwas_httphead",
        importpath = "github.com/gobwas/httphead",
        sum = "h1:s+21KNqlpePfkah2I+gwHF8xmJWRjooY+5248k6m4A0=",
        version = "v0.0.0-20180130184737-2c6c146eadee",
    )
    go_repository(
        name = "com_github_gobwas_pool",
        importpath = "github.com/gobwas/pool",
        sum = "h1:QEmUOlnSjWtnpRGHF3SauEiOsy82Cup83Vf2LcMlnc8=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_gobwas_ws",
        importpath = "github.com/gobwas/ws",
        sum = "h1:CoAavW/wd/kulfZmSIBt6p24n4j7tHgNVCjsfHVNUbo=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )

    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
        version = "v0.0.0-20160126235308-23def4e6c14b",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt",
        importpath = "github.com/golang-jwt/jwt",
        sum = "h1:IfV12K8xAKAnZqdXVzCZ+TOjboZ2keLg81eXfW3O+oY=",
        version = "v3.2.2+incompatible",
    )

    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_gomodule_redigo",
        importpath = "github.com/gomodule/redigo",
        sum = "h1:y0Wmhvml7cGnzPa9nocn/fMraMH/lMDdeG+rkx4VgYY=",
        version = "v1.7.1-0.20190724094224-574c33c3df38",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:81/ik6ipDQS2aGcBfIN5dHDB36BwrStyeAQquSYCV4o=",
        version = "v0.5.7",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        importpath = "github.com/google/go-querystring",
        sum = "h1:Xkwi/a1rcvNg1PPYe5vI8GbeBY/jrVuDX5ASuANWTrk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:EVhdT+1Kseyi1/pUmXKaFxYsDNy9RQYkMWRH68J/W7Y=",
        version = "v1.1.2",
    )

    go_repository(
        name = "com_github_gopherjs_gopherjs",
        importpath = "github.com/gopherjs/gopherjs",
        sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
        version = "v0.0.0-20181017120253-0766667cb4d1",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:q7AeDBpnBk8AogcD4DSag/Ukw/KV+YhzLj2bP5HvKCM=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
        version = "v1.16.0",
    )

    go_repository(
        name = "com_github_hashicorp_go_version",
        importpath = "github.com/hashicorp/go-version",
        sum = "h1:3vNe/fWF5CBgRIguda1meWhsZHy3m8gCJ5wx+dIzX/E=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_imkira_go_interpol",
        importpath = "github.com/imkira/go-interpol",
        sum = "h1:KIiKr0VSG2CUW1hl1jpiyuzuJeKUUpC8iM1AIE7N1Vk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        importpath = "github.com/inconshreveable/mousetrap",
        sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_iris_contrib_blackfriday",
        importpath = "github.com/iris-contrib/blackfriday",
        sum = "h1:o5sHQHHm0ToHUlAJSTjW9UWicjJSDDauOOQ2AHuIVp4=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_iris_contrib_go_uuid",
        importpath = "github.com/iris-contrib/go.uuid",
        sum = "h1:XZubAYg61/JwnJNbZilGjf3b3pB80+OQg2qf6c8BfWE=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_iris_contrib_i18n",
        importpath = "github.com/iris-contrib/i18n",
        sum = "h1:Kyp9KiXwsyZRTeoNjgVCrWks7D8ht9+kg6yCjh8K97o=",
        version = "v0.0.0-20171121225848-987a633949d0",
    )
    go_repository(
        name = "com_github_iris_contrib_jade",
        importpath = "github.com/iris-contrib/jade",
        sum = "h1:p7J/50I0cjo0wq/VWVCDFd8taPJbuFC+bq23SniRFX0=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_iris_contrib_pongo2",
        importpath = "github.com/iris-contrib/pongo2",
        sum = "h1:zGP7pW51oi5eQZMIlGA3I+FHY9/HOQWDB+572yin0to=",
        version = "v0.0.1",
    )

    go_repository(
        name = "com_github_iris_contrib_schema",
        importpath = "github.com/iris-contrib/schema",
        sum = "h1:10g/WnoRR+U+XXHWKBHeNy/+tZmM2kcAVGLOsz+yaDA=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath_internal_testify",
        importpath = "github.com/jmespath/go-jmespath/internal/testify",
        sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_joker_hpp",
        importpath = "github.com/Joker/hpp",
        sum = "h1:65+iuJYdRXv/XyN62C1uEmmOx3432rNG/rKlX6V7Kkc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_joker_jade",
        importpath = "github.com/Joker/jade",
        sum = "h1:mreN1m/5VJ/Zc3b4pzj9qU6D9SRQ6Vm+3KfI328t3S8=",
        version = "v1.0.1-0.20190614124447-d475f43051e7",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:9yzud/Ht36ygwatGx56VwCZtlI/2AD15T1X2sjSuGns=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_jtolds_gls",
        importpath = "github.com/jtolds/gls",
        sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
        version = "v4.20.0+incompatible",
    )
    go_repository(
        name = "com_github_juju_errors",
        importpath = "github.com/juju/errors",
        sum = "h1:rhqTjzJlm7EbkELJDKMTU7udov+Se0xZkWmugr6zGok=",
        version = "v0.0.0-20181118221551-089d3ea4e4d5",
    )
    go_repository(
        name = "com_github_juju_loggo",
        importpath = "github.com/juju/loggo",
        sum = "h1:MK144iBQF9hTSwBW/9eJm034bVoG30IshVm688T2hi8=",
        version = "v0.0.0-20180524022052-584905176618",
    )
    go_repository(
        name = "com_github_juju_testing",
        importpath = "github.com/juju/testing",
        sum = "h1:WQM1NildKThwdP7qWrNAFGzp4ijNLw8RlgENkaI4MJs=",
        version = "v0.0.0-20180920084828-472a3e8b2073",
    )
    go_repository(
        name = "com_github_k0kubun_colorstring",
        importpath = "github.com/k0kubun/colorstring",
        sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
        version = "v0.0.0-20150214042306-9440f1994b88",
    )
    go_repository(
        name = "com_github_kataras_golog",
        importpath = "github.com/kataras/golog",
        sum = "h1:vRDRUmwacco/pmBAm8geLn8rHEdc+9Z4NAr5Sh7TG/4=",
        version = "v0.0.10",
    )
    go_repository(
        name = "com_github_kataras_iris_v12",
        importpath = "github.com/kataras/iris/v12",
        sum = "h1:O3gJasjm7ZxpxwTH8tApZsvf274scSGQAUpNe47c37U=",
        version = "v12.1.8",
    )
    go_repository(
        name = "com_github_kataras_neffos",
        importpath = "github.com/kataras/neffos",
        sum = "h1:pdJaTvUG3NQfeMbbVCI8JT2T5goPldyyfUB2PJfh1Bs=",
        version = "v0.0.14",
    )
    go_repository(
        name = "com_github_kataras_pio",
        importpath = "github.com/kataras/pio",
        sum = "h1:6NAi+uPJ/Zuid6mrAKlgpbI11/zK/lV4B2rxWaJN98Y=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_kataras_sitemap",
        importpath = "github.com/kataras/sitemap",
        sum = "h1:4HCONX5RLgVy6G4RkYOV3vKNcma9p236LdGOipJsaFE=",
        version = "v0.0.5",
    )

    go_repository(
        name = "com_github_kisielk_errcheck",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:xqfchp4whNFxn5A4XFyyYtitiWI8Hy5EW59jEwcyL6U=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_klauspost_cpuid",
        importpath = "github.com/klauspost/cpuid",
        sum = "h1:vJi+O/nMdFt0vqm8NZBI6wzALWdA2X+egi0ogNyrC/w=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:mweAR1A6xJ3oS2pRaGiHgQ4OO8tzTaLawm8vnODuwDk=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_kr_logfmt",
        importpath = "github.com/kr/logfmt",
        sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
        version = "v0.0.0-20140226030751-b84e30acd515",
    )

    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_labstack_echo_v4",
        importpath = "github.com/labstack/echo/v4",
        sum = "h1:JXk6H5PAw9I3GwizqUHhYyS4f45iyGebR/c1xNCeOCY=",
        version = "v4.5.0",
    )
    go_repository(
        name = "com_github_labstack_gommon",
        importpath = "github.com/labstack/gommon",
        sum = "h1:JEeO0bvc78PKdyHxloTKiF8BD5iGrH8T6MSeGvSgob0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_leodido_go_urn",
        importpath = "github.com/leodido/go-urn",
        sum = "h1:hpXL4XnriNwQ/ABnpepYM/1vCLWNDfUNts8dX3xTG6Y=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:LLgXmsheXeRoUOBOjtwPQCWIYqM/LU1ayDtDePerRcY=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:nQ+aFkoE2TMGc0b68U2OKSexC+eq46+XwZzWXHRmPYs=",
        version = "v0.1.11",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:yVuAays6BHfxijgZPzw+3Zlu5yQgKGP2/hcQbHb7S9Y=",
        version = "v0.0.14",
    )
    go_repository(
        name = "com_github_mattn_goveralls",
        importpath = "github.com/mattn/goveralls",
        sum = "h1:7eJB6EqsPhRVxvwEXGnqdO2sJI0PTsrWoTMXEk9/OQc=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_mediocregopher_mediocre_go_lib",
        importpath = "github.com/mediocregopher/mediocre-go-lib",
        sum = "h1:3dQJqqDouawQgl3gBE1PNHKFkJYGEuFb1DbSlaxdosE=",
        version = "v0.0.0-20181029021733-cb65787f37ed",
    )
    go_repository(
        name = "com_github_mediocregopher_radix_v3",
        importpath = "github.com/mediocregopher/radix/v3",
        sum = "h1:galbPBjIwmyREgwGCfQEN4X8lxbJnKBYurgz+VfcStA=",
        version = "v3.4.2",
    )
    go_repository(
        name = "com_github_microcosm_cc_bluemonday",
        importpath = "github.com/microcosm-cc/bluemonday",
        sum = "h1:5lPfLTTAvAbtS0VqT+94yOtFnGfUWYyx0+iToC3Os3s=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_moul_http2curl",
        importpath = "github.com/moul/http2curl",
        sum = "h1:dRMWoAtb+ePxMlLkrCbAqh4TlPHXvoGUSQ323/9Zahs=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_nats_io_jwt",
        importpath = "github.com/nats-io/jwt",
        sum = "h1:xdnzwFETV++jNc4W1mw//qFyJGb2ABOombmZJQS4+Qo=",
        version = "v0.3.0",
    )

    go_repository(
        name = "com_github_nats_io_nats_go",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:ik3HbLhZ0YABLto7iX80pZLPw/6dx3T+++MZJwLnMrQ=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:qMd4+pRHgdr1nAClu+2h/2a5F2TmKcCzjCDazVgRoX4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_oneofone_xxhash",
        importpath = "github.com/OneOfOne/xxhash",
        sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
        version = "v1.2.2",
    )

    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:OoxbjfXVZyod1fmWYhI7SEyaD8B00ynP3T+D5GiyHOY=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:K0jcRCwNQM3vFGh1ppMtDh/+7ApJrjldlX8fA0jDTLQ=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_pelletier_go_toml",
        importpath = "github.com/pelletier/go-toml",
        sum = "h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_pingcap_errors",
        importpath = "github.com/pingcap/errors",
        sum = "h1:lFuQV/oaUMGcD2tqt+01ROSmJs75VG1ToEOkZIZ4nE4=",
        version = "v0.11.4",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version = "v0.0.0-20190812154241-14fe0d1b01d4",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_russross_blackfriday_v2",
        importpath = "github.com/russross/blackfriday/v2",
        sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_ryanuber_columnize",
        importpath = "github.com/ryanuber/columnize",
        sum = "h1:j1Wcmh8OrK4Q7GXY+V7SVSY8nUWQxHW5TkBe7YUl+2s=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_schollz_closestmatch",
        importpath = "github.com/schollz/closestmatch",
        sum = "h1:Uel2GXEpJqOWBrlyI+oY9LTiyyjYS17cCYRqP13/SHk=",
        version = "v2.1.0+incompatible",
    )

    go_repository(
        name = "com_github_sergi_go_diff",
        importpath = "github.com/sergi/go-diff",
        sum = "h1:Kpca3qRNrduNnOQeazBd0ysaKrUJiIuISHxogkT9RPQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_shopify_goreferrer",
        importpath = "github.com/Shopify/goreferrer",
        sum = "h1:WDC6ySpJzbxGWFh4aMxFFC28wwGp5pEuoTtvA4q/qQ4=",
        version = "v0.0.0-20181106222321-ec9c9a553398",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:SPIRibHv4MatM3XXNO2BJeFLZwZ2LvZgfQ5+UNI2im4=",
        version = "v1.4.2",
    )

    go_repository(
        name = "com_github_smartystreets_assertions",
        importpath = "github.com/smartystreets/assertions",
        sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
        version = "v0.0.0-20180927180507-b2de0cb4f26d",
    )
    go_repository(
        name = "com_github_smartystreets_goconvey",
        importpath = "github.com/smartystreets/goconvey",
        sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_github_spaolacci_murmur3",
        importpath = "github.com/spaolacci/murmur3",
        sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
        version = "v0.0.0-20180118202830-f09979ecbc72",
    )

    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:m8/z1t7/fwjysjQRYbP0RD+bUIF/8tJwPdEZsI83ACI=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        sum = "h1:f0B+LkLX6DtmRH1isoNA9VTtNUK9K8xYd28JNNfOv/s=",
        version = "v0.0.5",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:zPAT6CGy6wXeQ7NtTnaTerfKOsV6V6F8agHXFiazDkg=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        sum = "h1:VUFqw5KcqRf7i70GOzW7N+Q7+gxVBkSSqiXB12+JQ4M=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:2vfRuCMp5sSVIDSqO8oNnWJq7mPa6KVP3iPIwFBuy8A=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_ugorji_go",
        importpath = "github.com/ugorji/go",
        sum = "h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_ugorji_go_codec",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:2SvQaVZ1ouYrrKKwoSk2pzd4A9evlKJb9oTL+OaLUSs=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_urfave_cli_v2",
        importpath = "github.com/urfave/cli/v2",
        sum = "h1:JTTnM6wKzdA0Jqodd966MVj4vWbbquZykeX1sKbe2C4=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_urfave_negroni",
        importpath = "github.com/urfave/negroni",
        sum = "h1:kIimOitoypq34K7TG7DUaJ9kq/N4Ofuwi1sjz0KipXc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_bytebufferpool",
        importpath = "github.com/valyala/bytebufferpool",
        sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_fasthttp",
        importpath = "github.com/valyala/fasthttp",
        sum = "h1:d3AAQJ2DRcxJYHm7OXNXtXt2as1vMDfxeIcFvhmGGm4=",
        version = "v1.34.0",
    )
    go_repository(
        name = "com_github_valyala_fasttemplate",
        importpath = "github.com/valyala/fasttemplate",
        sum = "h1:TVEnxayobAdVkhQfrfes2IzOB6o+z4roRkPF52WA1u4=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_valyala_tcplisten",
        importpath = "github.com/valyala/tcplisten",
        sum = "h1:rBHj/Xf+E1tRGZyWIWwJDiRY0zc1Js+CV5DqwacVSA8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
        version = "v0.0.0-20180127040702-4e3ac2762d5f",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xordataexchange_crypt",
        importpath = "github.com/xordataexchange/crypt",
        sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
        version = "v0.0.3-0.20170626215501-b2862e3d0a77",
    )
    go_repository(
        name = "com_github_yalp_jsonpath",
        importpath = "github.com/yalp/jsonpath",
        sum = "h1:6fRhSjgLCkTD3JnJxvaJ4Sj+TYblw757bqYgZaOq5ZY=",
        version = "v0.0.0-20180802001716-5cc68e5049a0",
    )
    go_repository(
        name = "com_github_yudai_gojsondiff",
        importpath = "github.com/yudai/gojsondiff",
        sum = "h1:27cbfqXLVEJ1o8I6v3y9lg8Ydm53EKqHXAOMxEGlCOA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_yudai_golcs",
        importpath = "github.com/yudai/golcs",
        sum = "h1:BHyfKlQyqbsFN5p3IfnEUduWvb9is428/nNb5L3U01M=",
        version = "v0.0.0-20170316035057-ecda9a501e82",
    )
    go_repository(
        name = "com_github_yudai_pp",
        importpath = "github.com/yudai/pp",
        sum = "h1:Q4//iY4pNF6yPLZIigmvcl7k/bPgrcTPIFIcmawg5bI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:dPmz1Snjq0kmkz159iL7S6WzdahUTHnHB5M56WFVifs=",
        version = "v1.3.5",
    )

    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:eOI3/cP2VTU6uZLDYAoic+eyzzB9YyGmJ7eIjl8rOPg=",
        version = "v0.34.0",
    )
    go_repository(
        name = "com_shuralyov_dmitri_gpu_mtl",
        importpath = "dmitri.shuralyov.com/gpu/mtl",
        sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
        version = "v0.0.0-20190408044501-666a987793e9",
    )

    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:qIbj1fsPNlZgppZ+VLlY7N33q108Sa+fhmuc+sWQYwY=",
        version = "v1.0.0-20180628173108-788fd7840127",
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_go_playground_assert_v1",
        importpath = "gopkg.in/go-playground/assert.v1",
        sum = "h1:xoYuJVE7KT85PYWrN730RguIQO0ePzVRfFMXadIrXTM=",
        version = "v1.2.1",
    )
    go_repository(
        name = "in_gopkg_go_playground_validator_v8",
        importpath = "gopkg.in/go-playground/validator.v8",
        sum = "h1:lFB4DoMU6B626w8ny76MV7VX6W2VHct2GVOI3xgiMrQ=",
        version = "v8.18.2",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:GyboHr4UqMiLUybYjd22ZjQIKEJEpgtLXtuGbR21Oho=",
        version = "v1.51.1",
    )

    go_repository(
        name = "in_gopkg_mgo_v2",
        importpath = "gopkg.in/mgo.v2",
        sum = "h1:xcEWjVhvbDy+nHP67nPDDpbYrY+ILlfndk4bRioVHaU=",
        version = "v2.0.0-20180705113604-9856a29383ce",
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:h8qDotaEPuJATrMmW04NCwg7v22aHH28wwpauUhK9Oo=",
        version = "v3.0.0-20210107192922-496545a6307b",
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        importpath = "sigs.k8s.io/yaml",
        sum = "h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
        version = "v0.7.0",
    )

    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
        version = "v1.4.0",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:YxHp5zqIcAShDEvRr5/0rVESVS+njYF68PSdazrNLJo=",
        version = "v0.0.0-20220126215142-9970aeb2e350",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:weqSxi/TMs1SqFRMHCtBgXRs8k3X39QIDEZ0pRcttUg=",
        version = "v1.44.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:w43yiav+6bVFTBQFZX0r7ipe9JQ1QsbMgHwbBziscLw=",
        version = "v1.28.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:f+lwQ+GtmgoY+A2YaQxlSOnDjXcQ7ZRLWOHbC6HtRqE=",
        version = "v0.0.0-20220214200702-86341886e292",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
        version = "v0.0.0-20190121172915-509febef88a4",
    )
    go_repository(
        name = "org_golang_x_image",
        importpath = "golang.org/x/image",
        sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
        version = "v0.0.0-20190802002840-cff245a6509b",
    )

    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:5hukYrvBGR8/eNkX5mdUezrA6JiaEZDtJb9Ei+1LlBs=",
        version = "v0.0.0-20190930215403-16217165b5de",
    )
    go_repository(
        name = "org_golang_x_mobile",
        importpath = "golang.org/x/mobile",
        sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
        version = "v0.0.0-20190719004257-d2bd2a29d028",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:Gz96sIWK3OalVv/I/qNygP42zyoKp3xptRVCWRFEBvo=",
        version = "v0.4.2",
    )

    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:oA4XRj0qtSt8Yo1Zms0CUlsT3KG69V2UGQWPBxujDmc=",
        version = "v0.0.0-20220225172249-27dd8689420f",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:TzXSXBo42m9gQenoE3b9BGiEpg5IG2JkU5FkPIawgtw=",
        version = "v0.0.0-20200107190931-bf48bf16ab8d",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=",
        version = "v0.0.0-20210220032951-036812b2e83c",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:nhht2DYV/Sn3qOayu8lM+cU1ii9sTLUeBQwQQfUHtrs=",
        version = "v0.0.0-20220227234510-4e6760a101f9",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:JGgROgKl9N8DuW20oFS5gxc+lE67/N3FcwmBPMe7ArY=",
        version = "v0.0.0-20210927222741-03fcf44c2211",
    )

    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
        version = "v0.3.7",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:Hir2P/De0WpUhtrKGGjvSb2YxUgyZ7EFOSLIcSSpiwE=",
        version = "v0.0.0-20201208040808-7e3f01d25324",
    )

    go_repository(
        name = "org_golang_x_tools",
        build_extra_args = [
            "-exclude=**/testdata",
            "-exclude=go/packages/packagestest",
        ],
        importpath = "golang.org/x/tools",
        sum = "h1:ouewzE6p+/VEB31YYnTbEJdi8pFqKp4P4n85vwo3DHA=",
        version = "v0.1.5",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
        version = "v0.0.0-20200804184101-5ec99f83aff1",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_goleak",
        importpath = "go.uber.org/goleak",
        sum = "h1:wy28qYRKZgnJTxGxvye5/wgWr1EKjmUDGYox5mGlRlI=",
        version = "v1.1.11",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:zaiO/rmgFjbmCXdSYJWQcdvOCsthmdaHfr3Gm2Kx4Ec=",
        version = "v1.7.0",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:WefMeulhovoZ2sYXz7st6K0sLj7bBhpiFaud4r4zST8=",
        version = "v1.21.0",
    )

def go_dependenciest():
    pass
