package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `<td>A1100015</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1100015' target='_blank'>形势与政策</a></td>
                <td >1</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考查</td>
                <td >1,4,5,7</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1100022</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1100022' target='_blank'>思想道德与法治</a></td>
                <td >3</td>
                <td >48</td><td>0</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >1</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1100041</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1100041' target='_blank'>中国近现代史纲要</a></td>
                <td >3</td>
                <td >48</td><td>0</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >2</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1100061</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1100061' target='_blank'>思想政治理论课实践教学</a></td>
                <td >2</td>
                <td >0</td><td>48</td>
                <td ></td><td >48</td>
                <td>考查</td>
                <td >2</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1100032</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1100032' target='_blank'>马克思主义基本原理</a></td>
                <td >3</td>
                <td >48</td><td>0</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >3</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1100051</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1100051' target='_blank'>毛泽东思想和中国特色社会主义理论体系概论</a></td>
                <td >4</td>
                <td >64</td><td>0</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >4</td>
                <td>必修</td>
                </tr><tr><td align='center'><b>学分要求</b></td><td colspan='9'>必修：16.0 学分，选修：.0 学分</td></tr><tr><td align='center'><b>备注</b></td><td colspan='9'>思想政治理论课实践教学第2学期安排在课表，第3-6学期在网络教学平台提交规定材料。</td></tr></tbody>
    </table>（2）体育与军事 <table  class='pTable'>
    <thead> 
    <tr><td rowspan='2'>课程号</td><td rowspan='2'>课程名称</td><td rowspan='2'>学分</td>
    <td colspan='4'>学时</td><td rowspan='2'>考核方式</td><td rowspan='2'>开课学期</td><td rowspan='2'>修读要求</td></tr>
    <tr><td>理论</td><td>实验</td><td>实践</td><td>合计</td></tr>
    </thead>
    <tbody><tr align='center'>
                <td>A1090010</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1090010' target='_blank'>大学体育1</a></td>
                <td >1</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >1</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1000010</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1000010' target='_blank'>军事技能</a></td>
                <td >2</td>
                <td >0</td><td>0</td>
                <td >2.0周</td><td >2.0周</td>
                <td>考查</td>
                <td >1</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1170010</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1170010' target='_blank'>军事理论</a></td>
                <td >2</td>
                <td >36</td><td>0</td>
                <td ></td><td >36</td>
                <td>考试</td>
                <td >1</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1090020</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1090020' target='_blank'>大学体育2</a></td>
                <td >1</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >2</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1090031</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1090031' target='_blank'>体育（俱乐部）</a></td>
                <td >2</td>
                <td >80</td><td>0</td>
                <td ></td><td >80</td>
                <td>考查</td>
                <td >3,4,5,6,7</td>
                <td>必修</td>
                </tr><tr><td align='center'><b>学分要求</b></td><td colspan='9'>必修：8.0 学分，选修：.0 学分</td></tr><tr><td align='center'><b>备注</b></td><td colspan='9'>	体育（俱乐部）学生在3-7学期选择4个学期修读，每学期20学时0.5学分。</td></tr></tbody>
    </table>（3）大学外语 <table  class='pTable'>
    <thead> 
    <tr><td rowspan='2'>课程号</td><td rowspan='2'>课程名称</td><td rowspan='2'>学分</td>
    <td colspan='4'>学时</td><td rowspan='2'>考核方式</td><td rowspan='2'>开课学期</td><td rowspan='2'>修读要求</td></tr>
    <tr><td>理论</td><td>实验</td><td>实践</td><td>合计</td></tr>
    </thead>
    <tbody><tr align='center'>
                <td>A1050041</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050041' target='_blank'>大学英语2</a></td>
                <td >4</td>
                <td >48</td><td>16</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >1,2</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1050031</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050031' target='_blank'>大学英语1</a></td>
                <td >4</td>
                <td >48</td><td>16</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >1</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1050051</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050051' target='_blank'>大学英语3</a></td>
                <td >4</td>
                <td >48</td><td>16</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >2</td>
                <td>必修</td>
                </tr><tr align='center'>
                <td>A1050090</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050090' target='_blank'>影视英语与文化</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >3</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050230</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050230' target='_blank'>英语口语与演讲</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >3</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050061</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050061' target='_blank'>综合英语</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >3</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050110</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050110' target='_blank'>英语国家社会与文化</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >4</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050100</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050100' target='_blank'>中国文化传播</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >4</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050140</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050140' target='_blank'>商务交际英语</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >3</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050120</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050120' target='_blank'>国际职场文化</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >3</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050240</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050240' target='_blank'>学术英语阅读与写作</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >4</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050270</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050270' target='_blank'>国际教育英语（雅思）</a></td>
                <td >4</td>
                <td >64</td><td>0</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >4</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050280</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050280' target='_blank'>国际教育英语（托福）</a></td>
                <td >4</td>
                <td >64</td><td>0</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >4</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050250</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050250' target='_blank'>科技英语阅读与翻译</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >4</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050260</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050260' target='_blank'>实用英语</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5,6</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050290</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050290' target='_blank'>俄语1</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050200</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050200' target='_blank'>法语1</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050180</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050180' target='_blank'>德语1</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050310</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050310' target='_blank'>西班牙语1</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050160</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050160' target='_blank'>日语1</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050170</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050170' target='_blank'>日语2</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >6</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050190</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050190' target='_blank'>德语2</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >6</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050210</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050210' target='_blank'>法语2</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >6</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050300</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050300' target='_blank'>俄语2</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >6</td>
                <td>选修</td>
                </tr><tr align='center'>
                <td>A1050320</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1050320' target='_blank'>西班牙语2</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >6</td>
                <td>选修</td>
                </tr><tr><td align='center'><b>学分要求</b></td><td colspan='9'>必修：8.0 学分，选修：4.0 学分</td></tr><tr><td align='center'><b>备注</b></td><td colspan='9'>大学英语1-3：根据学生入学英语水平实施分类、分级教学，必修8学分。</td></tr></tbody>
    </table>（4）计算机与大数据智能公共课程 <table  class='pTable'>
    <thead> 
    <tr><td rowspan='2'>课程号</td><td rowspan='2'>课程名称</td><td rowspan='2'>学分</td>
    <td colspan='4'>学时</td><td rowspan='2'>考核方式</td><td rowspan='2'>开课学期</td><td rowspan='2'>修读要求</td></tr>
    <tr><td>理论</td><td>实验</td><td>实践</td><td>合计</td></tr>
    </thead>
    <tbody><tr><td align='center'><b>学分要求</b></td><td colspan='9'>必修： 学分，选修： 学分</td></tr><tr><td align='center'><b>备注</b></td><td colspan='9'></td></tr></tbody>
    </table>（5）素质教育拓展课程<br> <table  class='pTable'>
<thead>  
<tr><td rowspan='2'  style='width:100px'>课程模块</td><td rowspan='2'>课程号</td><td rowspan='2'>课程名称</td><td rowspan='2'>学分</td>
<td colspan='4'>学时</td><td rowspan='2'>考核方式</td><td rowspan='2'>开课学期</td><td rowspan='2'>修读要求</td></tr>
<tr><td>理论</td><td>实验</td><td>实践</td><td>合计</td></tr>
</thead>
<tbody>
<tr><td colspan='10'>①素质拓展课程必修课</td></tr> <tr align='center'>
    <td rowspan='8'>必修模块</td> 
        <td>A1000020</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1000020' target='_blank'>校史校情和校纪校规教育</a></td>
        <td >0.5</td>
        <td >8</td><td>0</td>
        <td ></td><td >8</td>
        <td>考查</td>
        <td >1</td>
        <td>必修</td>
        </tr><tr align='center'> 
        <td>A1000030</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1000030' target='_blank'>新生研讨课/专业概论</a></td>
        <td >0.5</td>
        <td >8</td><td>0</td>
        <td ></td><td >8</td>
        <td>考查</td>
        <td >1</td>
        <td>必修</td>
        </tr><tr align='center'> 
        <td>A1170040</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1170040' target='_blank'>大学生职业发展与就业指导（1）</a></td>
        <td >1</td>
        <td >16</td><td>0</td>
        <td ></td><td >16</td>
        <td>考查</td>
        <td >1</td>
        <td>必修</td>
        </tr><tr align='center'> 
        <td>A1170020</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1170020' target='_blank'>大学生心理健康教育（1）</a></td>
        <td >1</td>
        <td >16</td><td>0</td>
        <td ></td><td >16</td>
        <td>考查</td>
        <td >2</td>
        <td>必修</td>
        </tr><tr align='center'> 
        <td>A1170030</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1170030' target='_blank'>大学生心理健康教育（2）</a></td>
        <td >1</td>
        <td >16</td><td>0</td>
        <td ></td><td >16</td>
        <td>考查</td>
        <td >5</td>
        <td>必修</td>
        </tr><tr align='center'> 
        <td>A1170050</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1170050' target='_blank'>大学生职业发展与就业指导（2）</a></td>
        <td >1</td>
        <td >16</td><td>0</td>
        <td ></td><td >16</td>
        <td>考查</td>
        <td >6</td>
        <td>必修</td>
         </tr><tr><td colspan=2>劳动教育</td><td colspan=8>融入学生日常学习和生活，各专业可安排参加实验室卫生和设备整理工作，寒暑假参加日常家务、学工学农、社会实践等，学生处组织学生打扫寝室卫生、整理内务和参加其他公益劳动来实施。</td></tr>
    <tr><td colspan=2>安全教育</td><td colspan=8>主要通过安全知识讲座和专项培训活动培养学生的自我保护能力及良好的应急心态，由保卫处牵头组织，学工系统组织实施。在军事理论、形式政策、安全教育工作中适当融入国防安全教育的内容。</td></tr><tr><td colspan='10'>②素质拓展其他类课程</td></tr><tr align='center'><td rowspan='6'>创新与创业</td> 
            <td>A1075230</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1075230' target='_blank'>大学生创业基础</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1125200</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125200' target='_blank'>职场心理学</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1175030</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1175030' target='_blank'>ICT产品创新开发</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc71  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc71  ' target='_blank'>大学生创业概论与实践</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1175040</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1175040' target='_blank'>企业管理基础</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc72  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc72  ' target='_blank'>商业计划书写作指导</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'><td rowspan='10'>人文与沟通</td> 
            <td>A1125010</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125010' target='_blank'>古诗词鉴赏</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc33  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc33  ' target='_blank'>中西文化与文学专题比较</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc65  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc65  ' target='_blank'>中国传统文化(Mooc)</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc47  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc47  ' target='_blank'>国学素养</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1075240</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1075240' target='_blank'>普通话规范与演讲</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1175080</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1175080' target='_blank'>公文写作</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1075220</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1075220' target='_blank'>应用文写作</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc66  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc66  ' target='_blank'>公共关系与人际交往能力(Mooc)</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc64  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc64  ' target='_blank'>商务礼仪(Mooc)</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1175090</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1175090' target='_blank'>西方哲学述评</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'><td rowspan='7'>审美与修养</td> 
            <td>A1125190</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125190' target='_blank'>大学生审美修养</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc17  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc17  ' target='_blank'>美学原理（Mooc）</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1125220</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125220' target='_blank'>中西方绘画艺术鉴赏</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1125230</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125230' target='_blank'>艺术设计作品赏析</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1125250</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125250' target='_blank'>非遗与民间艺术</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1125260</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125260' target='_blank'>摄影</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1125160</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1125160' target='_blank'>生活创意与家居设计</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'><td rowspan='7'>科学与技术</td> 
            <td>A1115021</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1115021' target='_blank'>数学文化</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1115041</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1115041' target='_blank'>应用统计学</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1065090</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1065090' target='_blank'>自然科学经典导引</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1005010</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1005010' target='_blank'>科技文献检索与论文写作</a></td>
            <td >1</td>
            <td >8</td><td>8</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1015021</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1015021' target='_blank'>现代通信技术</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1015040</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1015040' target='_blank'>数据分析与量化投资</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1085020</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1085020' target='_blank'>物联网技术概论</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'><td rowspan='10'>环境与社会</td> 
            <td>A1065080</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1065080' target='_blank'>环境与生态文明</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1065060</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1065060' target='_blank'>食品安全与人类健康</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>Mooc79  </td><td align='left'><a href='http://cc.cqupt.edu.cn/Mooc79  ' target='_blank'>沟通心理学</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1075270</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1075270' target='_blank'>网络安全教育</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1035100</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1035100' target='_blank'>工程伦理</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考试</td>
            <td >学期滚动开出</td>
            <td>必修</td>
            </tr><tr align='center'> 
            <td>A1035120</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1035120' target='_blank'>工程管理与经济决策</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考试</td>
            <td >学期滚动开出</td>
            <td>必修</td>
            </tr><tr align='center'> 
            <td>A1075140</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1075140' target='_blank'>知识产权基础</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1035160</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1035160' target='_blank'>投资与理财</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1075290</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1075290' target='_blank'>网络安全法理论与实务</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr align='center'> 
            <td>A1075300</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1075300' target='_blank'>网络安全攻防</a></td>
            <td >1</td>
            <td >16</td><td>0</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >学期滚动开出</td>
            <td>选修</td>
            </tr><tr><td align='center'><b>学分要求</b></td><td colspan='10'>必修：9.0 学分，选修：3.0 学分</td></tr><tr><td align='center'><b>备注</b></td><td colspan='10'>除了《工程伦理》、《工程管理与经济决策》两门课程必修以外，素质教育拓展其他类课程选修总学分不低于3学分（MOOC最多选修2学分），其中审美与修养课程模块至少选修1学分。</td></tr></tbody>
</table><b>2、学科专业教育</b><br>（1）学科基础课程<br> <table  class='pTable'>
<thead>  
<tr><td rowspan='2'  style='width:100px'>课程模块</td><td rowspan='2'>课程号</td><td rowspan='2'>课程名称</td><td rowspan='2'>学分</td>
<td colspan='4'>学时</td><td rowspan='2'>考核方式</td><td rowspan='2'>开课学期</td><td rowspan='2'>修读要求</td></tr>
<tr><td>理论</td><td>实验</td><td>实践</td><td>合计</td></tr>
</thead>
<tbody><tr align='center'><td rowspan='4'>数学基础课程Ⅰ</td> 
            <td>A1110011</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110011' target='_blank'>高等数学A(上)</a></td>
            <td >5</td>
            <td >80</td><td>0</td>
            <td ></td><td >80</td>
            <td>考试</td>
            <td >1</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A1110032</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110032' target='_blank'>线性代数</a></td>
            <td >3</td>
            <td >48</td><td>0</td>
            <td ></td><td >48</td>
            <td>考试</td>
            <td >1</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A1110022</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110022' target='_blank'>高等数学A(下)</a></td>
            <td >6</td>
            <td >96</td><td>0</td>
            <td ></td><td >96</td>
            <td>考试</td>
            <td >2</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A1110140</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110140' target='_blank'>概率论与数理统计</a></td>
            <td >3</td>
            <td >48</td><td>0</td>
            <td ></td><td >48</td>
            <td>考试</td>
            <td >3</td>
            <td>必修 </td>
            </tr><tr align='center'><td rowspan='1'>数学基础课程Ⅱ</td> 
            <td>A1110190</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110190' target='_blank'>离散数学A</a></td>
            <td >4</td>
            <td >64</td><td>0</td>
            <td ></td><td >64</td>
            <td>考试</td>
            <td >3</td>
            <td>必修 </td>
            </tr><tr align='center'><td rowspan='4'>大学物理类课程</td> 
            <td>A1110362</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110362' target='_blank'>大学物理C（上）</a></td>
            <td >3.5</td>
            <td >56</td><td>0</td>
            <td ></td><td >56</td>
            <td>考试</td>
            <td >2</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A1110340</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110340' target='_blank'>大学物理实验B（上）</a></td>
            <td >1</td>
            <td >0</td><td>16</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >2</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A1110341</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110341' target='_blank'>大学物理实验B（下）</a></td>
            <td >1</td>
            <td >0</td><td>16</td>
            <td ></td><td >16</td>
            <td>考查</td>
            <td >3</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A1110363</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1110363' target='_blank'>大学物理C（下）</a></td>
            <td >3.5</td>
            <td >56</td><td>0</td>
            <td ></td><td >56</td>
            <td>考试</td>
            <td >3</td>
            <td>必修 </td>
            </tr><tr align='center'><td rowspan='2'>数字与逻辑电路课程</td> 
            <td>A2020390</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2020390' target='_blank'>数字与逻辑电路基础</a></td>
            <td >3</td>
            <td >48</td><td>0</td>
            <td ></td><td >48</td>
            <td>考试</td>
            <td >3</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A2021060</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2021060' target='_blank'>数字电路与逻辑设计实验A</a></td>
            <td >2</td>
            <td >0</td><td>32</td>
            <td ></td><td >32</td>
            <td>考查</td>
            <td >3</td>
            <td>必修 </td>
            </tr><tr align='center'><td rowspan='4'>计算机基础课程</td> 
            <td>A1040020</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1040020' target='_blank'>计算机科学导论</a></td>
            <td >2</td>
            <td >32</td><td>0</td>
            <td ></td><td >32</td>
            <td>考试</td>
            <td >1</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A2040900</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040900' target='_blank'>C语言程序设计（上）</a></td>
            <td >2</td>
            <td >24</td><td>8</td>
            <td ></td><td >32</td>
            <td>考试</td>
            <td >1</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A2040910</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040910' target='_blank'>C语言程序设计（下）</a></td>
            <td >2</td>
            <td >24</td><td>8</td>
            <td ></td><td >32</td>
            <td>考试</td>
            <td >2</td>
            <td>必修 </td>
            </tr><tr align='center'> 
            <td>A1040110</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1040110' target='_blank'>Python程序设计</a></td>
            <td >2</td>
            <td >16</td><td>16</td>
            <td ></td><td >32</td>
            <td>考试</td>
            <td >2</td>
            <td>必修 </td>
            </tr><tr><td align='center'><b>学分要求</b></td><td colspan='10'>必修：43.0 学分，选修：.0 学分</td></tr><tr><td align='center'><b>备注</b></td><td colspan='10'></td></tr></tbody>
</table>（2）专业教育课程<br> <table  class='pTable'>
    <thead>  
    <tr><td rowspan='2'  style='width:100px'>课程模块</td><td rowspan='2'>课程号</td><td rowspan='2'>课程名称</td><td rowspan='2'>学分</td>
    <td colspan='4'>学时</td><td rowspan='2'>考核方式</td><td rowspan='2'>开课学期</td><td rowspan='2'>修读要求</td></tr>
    <tr><td>理论</td><td>实验</td><td>实践</td><td>合计</td></tr>
    </thead>
    <tbody><tr align='center'><td rowspan='8'>专业基础课</td> 
                <td>A2041130</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2041130' target='_blank'>数据结构A</a></td>
                <td >4</td>
                <td >48</td><td>16</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >3</td>
                <td>必修 ★</td>
                </tr><tr align='center'> 
                <td>A2040360</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040360' target='_blank'>面向对象程序设计-C++</a></td>
                <td >3</td>
                <td >32</td><td>16</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >3</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040370</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040370' target='_blank'>面向对象程序设计-JAVA</a></td>
                <td >3</td>
                <td >32</td><td>16</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >3</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040100</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040100' target='_blank'>数据库原理(双语)</a></td>
                <td >4</td>
                <td >48</td><td>16</td>
                <td ></td><td >64</td>
                <td>考试</td>
                <td >4</td>
                <td>必修 ★</td>
                </tr><tr align='center'> 
                <td>A2040060</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040060' target='_blank'>计算机组织与结构</a></td>
                <td >3.5</td>
                <td >40</td><td>16</td>
                <td ></td><td >56</td>
                <td>考试</td>
                <td >4</td>
                <td>必修 ★</td>
                </tr><tr align='center'> 
                <td>A2040511</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040511' target='_blank'>算法分析与设计C</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >4</td>
                <td>必修 ★</td>
                </tr><tr align='center'> 
                <td>A2040030</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040030' target='_blank'>操作系统</a></td>
                <td >3</td>
                <td >48</td><td>0</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >4</td>
                <td>必修 ★</td>
                </tr><tr align='center'> 
                <td>A2040050</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040050' target='_blank'>计算机网络B</a></td>
                <td >3</td>
                <td >48</td><td>0</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >5</td>
                <td>必修 ★</td>
                </tr><tr align='center'><td rowspan='12'>专业课</td> 
                <td>A2110180</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2110180' target='_blank'>计算方法</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >3</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040960</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040960' target='_blank'>机器学习</a></td>
                <td >3</td>
                <td >48</td><td>0</td>
                <td ></td><td >48</td>
                <td>考试</td>
                <td >4</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A1040120</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1040120' target='_blank'>大数据导论</a></td>
                <td >1</td>
                <td >16</td><td>0</td>
                <td ></td><td >16</td>
                <td>考试</td>
                <td >4</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040040</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040040' target='_blank'>汇编语言程序设计</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >4</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040090</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040090' target='_blank'>软件工程A</a></td>
                <td >2.5</td>
                <td >40</td><td>0</td>
                <td ></td><td >40</td>
                <td>考试</td>
                <td >5</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040291</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040291' target='_blank'>计算机图形学B</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040081</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040081' target='_blank'>嵌入式系统设计B</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >5</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A1040130</td><td align='left'><a href='http://cc.cqupt.edu.cn/A1040130' target='_blank'>人工智能概论</a></td>
                <td >1</td>
                <td >16</td><td>0</td>
                <td ></td><td >16</td>
                <td>考试</td>
                <td >5</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040020</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040020' target='_blank'>编译原理</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >6</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2041200</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2041200' target='_blank'>物联网技术</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >6</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040521</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040521' target='_blank'>网络安全B</a></td>
                <td >2</td>
                <td >32</td><td>0</td>
                <td ></td><td >32</td>
                <td>考试</td>
                <td >7</td>
                <td>选修 </td>
                </tr><tr align='center'> 
                <td>A2040750</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040750' target='_blank'>计算机与智能科学前沿（微课）</a></td>
                <td >1</td>
                <td >16</td><td>0</td>
                <td ></td><td >16</td>
                <td>考试</td>
                <td >7</td>
                <td>选修 </td>
                </tr><tr align='center'><td rowspan='12'>专业集中实践</td> 
                <td>A2040810</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040810' target='_blank'>Linux系统实践</a></td>
                <td >1</td>
                <td >0</td><td>0</td>
                <td >1.0周</td><td >1.0周</td>
                <td>考查</td>
                <td >3</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2040200</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040200' target='_blank'>课程设计(操作系统)</a></td>
                <td >1</td>
                <td >0</td><td>0</td>
                <td >1.0周</td><td >1.0周</td>
                <td>考查</td>
                <td >4</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2040940</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040940' target='_blank'>Web综合设计实践</a></td>
                <td >2</td>
                <td >0</td><td>0</td>
                <td >2.0周</td><td >2.0周</td>
                <td>考查</td>
                <td >4</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2041250</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2041250' target='_blank'>ModelArts机器学习实践</a></td>
                <td >2</td>
                <td >0</td><td>0</td>
                <td >2.0周</td><td >2.0周</td>
                <td>考查</td>
                <td >4</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2040650</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040650' target='_blank'>云计算与大数据综合实践</a></td>
                <td >2</td>
                <td >0</td><td>0</td>
                <td >2.0周</td><td >2.0周</td>
                <td>考查</td>
                <td >5</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2041040</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2041040' target='_blank'>深度学习综合实践</a></td>
                <td >2</td>
                <td >0</td><td>0</td>
                <td >2.0周</td><td >2.0周</td>
                <td>考查</td>
                <td >5</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2040300</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040300' target='_blank'>课程设计（计算机网络）</a></td>
                <td >1</td>
                <td >0</td><td>0</td>
                <td >1.0周</td><td >1.0周</td>
                <td>考查</td>
                <td >5</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2040950</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2040950' target='_blank'>大数据系统开发实践</a></td>
                <td >2</td>
                <td >0</td><td>0</td>
                <td >2.0周</td><td >2.0周</td>
                <td>考查</td>
                <td >6</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2041050</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2041050' target='_blank'>数据工程综合实践</a></td>
                <td >2</td>
                <td >0</td><td>0</td>
                <td >2.0周</td><td >2.0周</td>
                <td>考查</td>
                <td >6</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A2041260</td><td align='left'><a href='http://cc.cqupt.edu.cn/A2041260' target='_blank'>计算机软件能力测评</a></td>
                <td >6</td>
                <td >0</td><td>0</td>
                <td >6.0周</td><td >6.0周</td>
                <td>考查</td>
                <td >6</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A200005E</td><td align='left'><a href='http://cc.cqupt.edu.cn/A200005E' target='_blank'>毕业实习</a></td>
                <td >8</td>
                <td >0</td><td>0</td>
                <td >8.0周</td><td >8.0周</td>
                <td>考查</td>
                <td >7</td>
                <td>必修 </td>
                </tr><tr align='center'> 
                <td>A200006Z</td><td align='left'><a href='http://cc.cqupt.edu.cn/A200006Z' target='_blank'>毕业设计（论文）</a></td>
                <td >14</td>
                <td >0</td><td>0</td>
                <td >16.0周</td><td >16.0周</td>
                <td>考查</td>
                <td >8</td>
                <td>必修 </td>
                </tr><tr><td align='center'><b>学分要求</b></td><td colspan='10'>必修：65.5 学分，选修：10.5 学分</td></tr><tr><td align='center'><b>备注</b></td><td colspan='10'>《面向对象程序设计-C++》、《面向对象程序设计-JAVA》两门课程至少选修一门。</td></tr></tbody>`
	//解析
	ret := regexp.MustCompile(`target='_blank'>(?s:(.*?))</a>`)
	//提取
	res := ret.FindAllStringSubmatch(str, -1)
	//打印
	for _, one := range res {
		fmt.Println(one[1])
	}
}
