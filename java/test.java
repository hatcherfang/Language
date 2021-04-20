import java.util.List;
import com.alibaba.fastjson.JSONObject;


public class test{
    public static void main(String[] args){

        JSONObject bodyJson = JSONObject.parseObject("{\"role_id\":[1, 2]}");
        List<String> s = bodyJson.get["role_id"];
        System.out.println(s);
        System.out.println(test.class);
    }
}
