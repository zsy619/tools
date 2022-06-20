package xstrip

import (
	"fmt"
	"regexp"
	"testing"
)

func strip_tags(content string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(content, "")
}

func Test_strip_tags(t *testing.T) {
	str :=
		`
<html>

<head>
    <meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>河南省大中专学生智慧就业平台 -- 后台管理</title>
<meta name="theme-color" content="#2929ff">
<meta name="keywords" content="招聘，大学生，在校，兼职，实习">
<meta name="description" content="专业为大学生提供最新最全的招聘信息">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
<meta http-equiv="X-UA-Compatible" content="IE=EmulateIE7" />

<meta name="renderer" content="webkit">
<link rel="stylesheet" href="/static/new/plugins/layui/css/layui.css" />
<link rel="stylesheet" href="/static/AmazeUI2.7.2/css/amazeui.min.css" />
<link rel="stylesheet" href="/static/new/css/pagenation.css" />
<link rel="stylesheet" href="/static/new/css/table.css" />
<link rel="stylesheet" href="/static/new/css/schooltag.css" />
<script rel="stylesheet" src="/static/new/js/common/commonCss.js"></script>
<script type="text/javascript" src="/static/new/js/common/commonJs.js"></script>
<script type="text/javascript" src="/static/new/plugins/misc/countTo/jquery.countTo.js"></script>
<script type="text/javascript" src="/static/new/plugins/layui/layui.js"></script>
<script type="text/javascript" src="/static/new/plugins/misc/kindeditor/kindeditor-all-min.js"></script>
<script type="text/javascript" src="/static/new/plugins/misc/kindeditor/lang/zh-CN.js"></script>
<script type="text/javascript" src="/static/new/js/jquery.form.js"></script>
<script type="text/javascript" src="/static/new/js/vue-2.6.14/vue.min.js"></script>
<script type="text/javascript" src="/static/new/js/jquery.pagination.js"></script>
<script type="text/javascript" src="/static/new/js/utils.js"></script>
    
    <style type="text/css">
        .error {
            color: #ec1b15;
            display: none;
        }
    </style>
</head>

<body id="root">
    <div id="header">
    <div class="navbar">
        <div class="navbar-header">
            <a class="navbar-brand" href="/">
                <img src="/static/new/images/logo.png" />
            </a>
        </div>
        <div class="pull-right">
            <div class="dropdown clearfix">
                <div class="toggle-person-msg" type="button" class="dropdown-toggle" id="dropdownMenu" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    <a href="javascript:void(0)">
                        管理员<span class="caret"></span>
                    </a>
                </div>
                <ul class="dropdown-menu" aria-labelledby="dropdownMenu">
                    <li>
                        <a href="/admin/logout"><i class="im-exit"></i>退出</a>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
    <div id="sidebar">
    <ul id="sideNav" class="nav nav-pills nav-stacked">
        <li>
            <a href="/admin/Index">首页
                <i class="fa fa-home"></i>
            </a>
        </li>
        <li>
            <a href="#">用人单位管理
                <i class="fa fa-sitemap"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/EmployRecrView/ManagerCompany">用人单位</a>
                </li>
                <li>
                    <a href="/admin/EmployRecrView/ManagerCompanyBlack">黑名单</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="/admin/EmployRecrView/ManagerPositionCompany">职位管理
                <i class="fa fa-tree"></i>
            </a>
        </li>
        <li>
            <a href="#">数据统计
                <i class="fa fa-sitemap"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/admin/JobApplyRecord?tab_active=1">投递记录</a>
                </li>
                <li>
                    <a href="/admin/admin/JobApplyRecord?tab_active=2">面试记录</a>
                </li>
                <li>
                    <a href="/admin/admin/Statistics">数据统计</a>
                </li>
                <li>
                    <a href="/admin/SticTationView/Jobfair">招聘会统计列表</a>
                </li>
                <li>
                    <a href="/admin/SticTationView/Index">数据展示</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">学校管理
                <i class="fa fa-sitemap"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/SchoolManageView/EditSchool">添加学校</a>
                </li>
                <li>
                    <a href="/admin/SchoolManageView/Index">学校列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">学生管理
                <i class="fa fa-user-o"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/StudentManageView/Details">投递情况</a>
                </li>
                <li>
                    <a href="/admin/StudentManageView/Index">学生列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">双选会管理
                <i class="fa fa-telegram"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/JobfairView/JobfairManagerList">双选会列表</a>
                </li>
                <li>
                    <a href="/admin/JobfairView/JobfairMeet">展位会场模板</a>
                </li>

                <li>
                    <a href="/admin/JobfairView/JobfairZphSetting">参数设置</a>
                </li>
                <li>
                    <a href="/admin/JobfairView/JobfairPublish">发布双选会</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">账号权限管理
                <i class="fa fa-telegram"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/AdminManageView/AdminEdit">添加账号</a>
                </li>
                <li>
                    <a href="/admin/AdminManageView/Index">账号列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">广告管理
                <i class="fa fa-telegram"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/AdManageView/Edit">添加广告</a>
                </li>
                <li>
                    <a href="/admin/AdManageView/Index">广告列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">简历配置
                <i class="fa fa-cog"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/ResumesManageView/EditCase">添加案例</a>
                </li>
                <li>
                    <a href="/admin/ResumesManageView/ResumeCase">案例管理</a>
                </li>
                <li>
                    <a href="/admin/ResumesManageView/EditTips">添加小贴士</a>
                </li>
                <li>
                    <a href="/admin/ResumesManageView/ResumeTips">小贴士管理</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">空中宣讲管理
                <i class="fa fa-cog"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/AirKenoteView/Edit">添加空中宣讲</a>
                </li>
                <li>
                    <a href="/admin/AirKenoteView/Index">空中宣讲列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="/admin/EmployRecrView/CompanyJobfair">用人单位专场管理
                <i class="fa fa-commenting-o"></i>
            </a>
        </li>
        <li>
            <a href="#">校招信息
                <i class="fa fa-cog"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/ProjectEngineerView/Edit">添加校招</a>
                </li>
                <li>
                    <a href="/admin/ProjectEngineerView/Index">校招列表</a>
                </li>
                <li>
                    <a href="/admin/ProjectEngineerView/Communication">交流群</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="/admin/admin/paramConfig">系统参数配置
                <i class="fa fa-cog"></i>
            </a>
        </li>
        <li>
            <a href="#">微信公众号
                <i class="fa fa-weixin"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/weixin/WechatAdminView/WechatConfigure">用人单位</a>
                </li>
                <li>
                    <a href="/weixin/JobSeekerWechatAdminView/WechatConfigure">求职者</a>
                </li>
            </ul>
        </li>
    </ul>
</div>
    <div id="contentOuter">
        <style type="text/css">
    .poster_box {
         
        margin: 30px;
        color: #303030;
        height: 680px
    }

    .poster_div {
        width: 350px;
        margin-right: 28px;
        float: left;
    }

    .poster_title {
        font-size: 14px;
    }

    .poster_bg {
        margin-top: 10px;
        width: 350px;
        height: 212px;
        border: 1px solid #dcdcdc;
        margin-bottom: 6px;
    }

    .poster_tips {
        font-size: 12px;
    }
</style>
<div id="content">
    <div class="content-wrapper">
        
        <div class="heading">
    <ul id="crumb" class="breadcrumb">
        <li>
            <i class="fa fa-home"></i><a href="/">首页</a>
        </li>
        
        <li> <i class="fa fa-angle-right"></i>双选会管理</li>
        
        <li> <i class="fa fa-angle-right"></i>更多操作</li>
        
        <li> <i class="fa fa-angle-right"></i>双选会海报</li>
        
    </ul>
</div>
        <span id="jobfair_id" style="display: none;">4</span>
        <div class="outlet detail-msg-box show-margin">
            <div class="row">
                <div class="bg-white btn-top-operate">
                    <span class="btn btn-sm btn-default" onclick="window.location.href='JobfairManagerList';">返回</span>
                </div>
                <div class="tab">
                    <ul class="tab-title" id="screen-btn">
    
    <li class='title-tab ' onclick='location.href="./functionalServiceManager?jobfair_id=4"'>用人单位增值服务</li>
    
    <li class='title-tab ' onclick='location.href="./jobfairMoreOperations?jobfair_id=4"'>二维码下载</li>
    <li class='title-tab ' onclick='location.href="./reportingManager?jobfair_id=4"'>报到管理</li>
    <li class='title-tab active' onclick='location.href="./posterDownload?jobfair_id=4"'>双选会海报</li>
    <li class='title-tab ' onclick='location.href="./jobfairParticipant?jobfair_id=4"'>参会码</li>
    <li class='title-tab ' onclick='location.href="./participantNoCompanyManager?jobfair_id=4"'>参会信息</li>
    <li class='title-tab ' onclick='location.href="./jobfairStudentSignList?jobfair_id=4"'>学生报名&签到</li>
</ul>
                    <div class="tab-content">
                        <div class="tab-item active">
                            <div class="table-box poster_box">

                                <div class="poster_div">
                                    <span class="poster_title">海报模板1</span>
                                    <img class="poster_bg" src="/static/new/template/posterPreview1.png" alt="走丢了~">
                                    <span class="poster_tips">尺寸: 145cm&times;88cm &nbsp;&nbsp;&nbsp;&nbsp;分辨率: 150dpi </span>
                                    <div class="am-margin-top">
                                        <a class="btn btn-default" download href="/static/new/template/posterBgOne.jpg">下载空白模板</a>
                                        <a class="btn btn-success" onclick="getPlaceNum(1)">批量下载海报</a>
                                    </div>
                                </div>

                                <div class="poster_div">
                                    <span class="poster_title">海报模板2</span>
                                    <img class="poster_bg" src="/static/new/template/posterPreview2.png" alt="走丢了~">
                                    <span class="poster_tips">尺寸: 145cm&times;88cm &nbsp;&nbsp;&nbsp;&nbsp;分辨率: 150dpi </span>
                                    <div class="am-margin-top">
                                        <a class="btn btn-default" download href="/static/new/template/posterBgTwo.jpg">下载空白模板</a>
                                        <a class="btn btn-success" onclick="getPlaceNum(2)">批量下载海报</a>
                                    </div>
                                </div>

                                <div class="poster_div" style="margin-top: 30px;">
                                    <span class="poster_title">海报模板3</span>
                                    <img class="poster_bg" src="/static/new/template/posterPreview3.png" alt="走丢了~">
                                    <span class="poster_tips">尺寸: 145cm&times;88cm &nbsp;&nbsp;&nbsp;&nbsp;分辨率: 150dpi </span>
                                    <div class="am-margin-top">
                                        <a class="btn btn-default" download href="/static/new/template/posterBgThree.jpg">下载空白模板</a>
                                        <a class="btn btn-success" onclick="getPlaceNum(3)">批量下载海报</a>
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="clearfix"></div>
</div>
</div>
<span style="display: none;" id="zhaopinhuiid">4</span>
<span id="is_school_content"></span>
<div class="popup">
</div>
<script type="text/javascript">
    var jobfairId = '4';
</script>
<script type="text/javascript" src="/static/new/admin_js/jobfair/posterDownload.js"></script>
    </div>
    <div class="footer">
    <div class="container">
        <p class="text-center">技术支持：河南省教育网有限公司</p>
        <p class="text-center"> 企业信息查询链接：
            <a href="https://www.cods.org.cn" target="_blank">全国组织机构代码管理中心查询</a>
            <a href="https://www.creditchina.gov.cn" target="_blank">信用中国</a>
            <a href="http://www.gsxt.gov.cn/index.html" target="_blank">国家企业信用公示系统</a>
            <a href="https://www.qcc.com" target="_blank">企查查</a>
        </p>
    </div>
</div>
    
</body>

</html>`
	fmt.Println("", strip_tags(str))
}

func TestStripTags(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"", ""},
		{"Hello, World!", "Hello, World!"},
		{"foo&amp;bar", "foo&amp;bar"},
		{`Hello <a href="www.example.com/">World</a>!`, "Hello World!"},
		{"Foo <textarea>Bar</textarea> Baz", "Foo Bar Baz"},
		{"Foo <!-- Bar --> Baz", "Foo  Baz"},
		{"<", "<"},
		{"foo < bar", "foo < bar"},
		{`Foo<script type="text/javascript">alert(1337)</script>Bar`, "FooBar"},
		{`Foo<div title="1>2">Bar`, "FooBar"},
		{`I <3 Ponies!`, `I <3 Ponies!`},
		{`<script>foo()</script>`, ``},
	}

	for _, test := range tests {
		if got := StripTags(test.input); got != test.want {
			t.Errorf("%q: want %q, got %q", test.input, test.want, got)
		}
	}
}

func Test_StripTags(t *testing.T) {
	str :=
		`
<html>

<head>
    <meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>河南省大中专学生智慧就业平台 -- 后台管理</title>
<meta name="theme-color" content="#2929ff">
<meta name="keywords" content="招聘，大学生，在校，兼职，实习">
<meta name="description" content="专业为大学生提供最新最全的招聘信息">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
<meta http-equiv="X-UA-Compatible" content="IE=EmulateIE7" />

<meta name="renderer" content="webkit">
<link rel="stylesheet" href="/static/new/plugins/layui/css/layui.css" />
<link rel="stylesheet" href="/static/AmazeUI2.7.2/css/amazeui.min.css" />
<link rel="stylesheet" href="/static/new/css/pagenation.css" />
<link rel="stylesheet" href="/static/new/css/table.css" />
<link rel="stylesheet" href="/static/new/css/schooltag.css" />
<script rel="stylesheet" src="/static/new/js/common/commonCss.js"></script>
<script type="text/javascript" src="/static/new/js/common/commonJs.js"></script>
<script type="text/javascript" src="/static/new/plugins/misc/countTo/jquery.countTo.js"></script>
<script type="text/javascript" src="/static/new/plugins/layui/layui.js"></script>
<script type="text/javascript" src="/static/new/plugins/misc/kindeditor/kindeditor-all-min.js"></script>
<script type="text/javascript" src="/static/new/plugins/misc/kindeditor/lang/zh-CN.js"></script>
<script type="text/javascript" src="/static/new/js/jquery.form.js"></script>
<script type="text/javascript" src="/static/new/js/vue-2.6.14/vue.min.js"></script>
<script type="text/javascript" src="/static/new/js/jquery.pagination.js"></script>
<script type="text/javascript" src="/static/new/js/utils.js"></script>
    
    <style type="text/css">
        .error {
            color: #ec1b15;
            display: none;
        }
    </style>
</head>

<body id="root">
    <div id="header">
    <div class="navbar">
        <div class="navbar-header">
            <a class="navbar-brand" href="/">
                <img src="/static/new/images/logo.png" />
            </a>
        </div>
        <div class="pull-right">
            <div class="dropdown clearfix">
                <div class="toggle-person-msg" type="button" class="dropdown-toggle" id="dropdownMenu" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    <a href="javascript:void(0)">
                        管理员<span class="caret"></span>
                    </a>
                </div>
                <ul class="dropdown-menu" aria-labelledby="dropdownMenu">
                    <li>
                        <a href="/admin/logout"><i class="im-exit"></i>退出</a>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
    <div id="sidebar">
    <ul id="sideNav" class="nav nav-pills nav-stacked">
        <li>
            <a href="/admin/Index">首页
                <i class="fa fa-home"></i>
            </a>
        </li>
        <li>
            <a href="#">用人单位管理
                <i class="fa fa-sitemap"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/EmployRecrView/ManagerCompany">用人单位</a>
                </li>
                <li>
                    <a href="/admin/EmployRecrView/ManagerCompanyBlack">黑名单</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="/admin/EmployRecrView/ManagerPositionCompany">职位管理
                <i class="fa fa-tree"></i>
            </a>
        </li>
        <li>
            <a href="#">数据统计
                <i class="fa fa-sitemap"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/admin/JobApplyRecord?tab_active=1">投递记录</a>
                </li>
                <li>
                    <a href="/admin/admin/JobApplyRecord?tab_active=2">面试记录</a>
                </li>
                <li>
                    <a href="/admin/admin/Statistics">数据统计</a>
                </li>
                <li>
                    <a href="/admin/SticTationView/Jobfair">招聘会统计列表</a>
                </li>
                <li>
                    <a href="/admin/SticTationView/Index">数据展示</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">学校管理
                <i class="fa fa-sitemap"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/SchoolManageView/EditSchool">添加学校</a>
                </li>
                <li>
                    <a href="/admin/SchoolManageView/Index">学校列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">学生管理
                <i class="fa fa-user-o"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/StudentManageView/Details">投递情况</a>
                </li>
                <li>
                    <a href="/admin/StudentManageView/Index">学生列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">双选会管理
                <i class="fa fa-telegram"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/JobfairView/JobfairManagerList">双选会列表</a>
                </li>
                <li>
                    <a href="/admin/JobfairView/JobfairMeet">展位会场模板</a>
                </li>

                <li>
                    <a href="/admin/JobfairView/JobfairZphSetting">参数设置</a>
                </li>
                <li>
                    <a href="/admin/JobfairView/JobfairPublish">发布双选会</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">账号权限管理
                <i class="fa fa-telegram"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/AdminManageView/AdminEdit">添加账号</a>
                </li>
                <li>
                    <a href="/admin/AdminManageView/Index">账号列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">广告管理
                <i class="fa fa-telegram"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/AdManageView/Edit">添加广告</a>
                </li>
                <li>
                    <a href="/admin/AdManageView/Index">广告列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">简历配置
                <i class="fa fa-cog"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/ResumesManageView/EditCase">添加案例</a>
                </li>
                <li>
                    <a href="/admin/ResumesManageView/ResumeCase">案例管理</a>
                </li>
                <li>
                    <a href="/admin/ResumesManageView/EditTips">添加小贴士</a>
                </li>
                <li>
                    <a href="/admin/ResumesManageView/ResumeTips">小贴士管理</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="#">空中宣讲管理
                <i class="fa fa-cog"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/AirKenoteView/Edit">添加空中宣讲</a>
                </li>
                <li>
                    <a href="/admin/AirKenoteView/Index">空中宣讲列表</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="/admin/EmployRecrView/CompanyJobfair">用人单位专场管理
                <i class="fa fa-commenting-o"></i>
            </a>
        </li>
        <li>
            <a href="#">校招信息
                <i class="fa fa-cog"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/admin/ProjectEngineerView/Edit">添加校招</a>
                </li>
                <li>
                    <a href="/admin/ProjectEngineerView/Index">校招列表</a>
                </li>
                <li>
                    <a href="/admin/ProjectEngineerView/Communication">交流群</a>
                </li>
            </ul>
        </li>
        <li>
            <a href="/admin/admin/paramConfig">系统参数配置
                <i class="fa fa-cog"></i>
            </a>
        </li>
        <li>
            <a href="#">微信公众号
                <i class="fa fa-weixin"></i>
            </a>
            <ul class="nav sub">
                <li>
                    <a href="/weixin/WechatAdminView/WechatConfigure">用人单位</a>
                </li>
                <li>
                    <a href="/weixin/JobSeekerWechatAdminView/WechatConfigure">求职者</a>
                </li>
            </ul>
        </li>
    </ul>
</div>
    <div id="contentOuter">
        <style type="text/css">
    .poster_box {
         
        margin: 30px;
        color: #303030;
        height: 680px
    }

    .poster_div {
        width: 350px;
        margin-right: 28px;
        float: left;
    }

    .poster_title {
        font-size: 14px;
    }

    .poster_bg {
        margin-top: 10px;
        width: 350px;
        height: 212px;
        border: 1px solid #dcdcdc;
        margin-bottom: 6px;
    }

    .poster_tips {
        font-size: 12px;
    }
</style>
<div id="content">
    <div class="content-wrapper">
        
        <div class="heading">
    <ul id="crumb" class="breadcrumb">
        <li>
            <i class="fa fa-home"></i><a href="/">首页</a>
        </li>
        
        <li> <i class="fa fa-angle-right"></i>双选会管理</li>
        
        <li> <i class="fa fa-angle-right"></i>更多操作</li>
        
        <li> <i class="fa fa-angle-right"></i>双选会海报</li>
        
    </ul>
</div>
        <span id="jobfair_id" style="display: none;">4</span>
        <div class="outlet detail-msg-box show-margin">
            <div class="row">
                <div class="bg-white btn-top-operate">
                    <span class="btn btn-sm btn-default" onclick="window.location.href='JobfairManagerList';">返回</span>
                </div>
                <div class="tab">
                    <ul class="tab-title" id="screen-btn">
    
    <li class='title-tab ' onclick='location.href="./functionalServiceManager?jobfair_id=4"'>用人单位增值服务</li>
    
    <li class='title-tab ' onclick='location.href="./jobfairMoreOperations?jobfair_id=4"'>二维码下载</li>
    <li class='title-tab ' onclick='location.href="./reportingManager?jobfair_id=4"'>报到管理</li>
    <li class='title-tab active' onclick='location.href="./posterDownload?jobfair_id=4"'>双选会海报</li>
    <li class='title-tab ' onclick='location.href="./jobfairParticipant?jobfair_id=4"'>参会码</li>
    <li class='title-tab ' onclick='location.href="./participantNoCompanyManager?jobfair_id=4"'>参会信息</li>
    <li class='title-tab ' onclick='location.href="./jobfairStudentSignList?jobfair_id=4"'>学生报名&签到</li>
</ul>
                    <div class="tab-content">
                        <div class="tab-item active">
                            <div class="table-box poster_box">

                                <div class="poster_div">
                                    <span class="poster_title">海报模板1</span>
                                    <img class="poster_bg" src="/static/new/template/posterPreview1.png" alt="走丢了~">
                                    <span class="poster_tips">尺寸: 145cm&times;88cm &nbsp;&nbsp;&nbsp;&nbsp;分辨率: 150dpi </span>
                                    <div class="am-margin-top">
                                        <a class="btn btn-default" download href="/static/new/template/posterBgOne.jpg">下载空白模板</a>
                                        <a class="btn btn-success" onclick="getPlaceNum(1)">批量下载海报</a>
                                    </div>
                                </div>

                                <div class="poster_div">
                                    <span class="poster_title">海报模板2</span>
                                    <img class="poster_bg" src="/static/new/template/posterPreview2.png" alt="走丢了~">
                                    <span class="poster_tips">尺寸: 145cm&times;88cm &nbsp;&nbsp;&nbsp;&nbsp;分辨率: 150dpi </span>
                                    <div class="am-margin-top">
                                        <a class="btn btn-default" download href="/static/new/template/posterBgTwo.jpg">下载空白模板</a>
                                        <a class="btn btn-success" onclick="getPlaceNum(2)">批量下载海报</a>
                                    </div>
                                </div>

                                <div class="poster_div" style="margin-top: 30px;">
                                    <span class="poster_title">海报模板3</span>
                                    <img class="poster_bg" src="/static/new/template/posterPreview3.png" alt="走丢了~">
                                    <span class="poster_tips">尺寸: 145cm&times;88cm &nbsp;&nbsp;&nbsp;&nbsp;分辨率: 150dpi </span>
                                    <div class="am-margin-top">
                                        <a class="btn btn-default" download href="/static/new/template/posterBgThree.jpg">下载空白模板</a>
                                        <a class="btn btn-success" onclick="getPlaceNum(3)">批量下载海报</a>
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="clearfix"></div>
</div>
</div>
<span style="display: none;" id="zhaopinhuiid">4</span>
<span id="is_school_content"></span>
<div class="popup">
</div>
<script type="text/javascript">
    var jobfairId = '4';
</script>
<script type="text/javascript" src="/static/new/admin_js/jobfair/posterDownload.js"></script>
    </div>
    <div class="footer">
    <div class="container">
        <p class="text-center">技术支持：河南省教育网有限公司</p>
        <p class="text-center"> 企业信息查询链接：
            <a href="https://www.cods.org.cn" target="_blank">全国组织机构代码管理中心查询</a>
            <a href="https://www.creditchina.gov.cn" target="_blank">信用中国</a>
            <a href="http://www.gsxt.gov.cn/index.html" target="_blank">国家企业信用公示系统</a>
            <a href="https://www.qcc.com" target="_blank">企查查</a>
        </p>
    </div>
</div>
    
</body>

</html>`
	fmt.Println("", StripTags(str))
}
