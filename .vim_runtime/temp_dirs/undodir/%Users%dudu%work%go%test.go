Vim�UnDo� (ܤ L7L�L�H�l(�-w�]Wj���-                    V       V   V   V    XB�j    _�                       	    ����                                                                                                                                                                                                                                                                                                                               	          
       v   
    XBrr     �               func Pop(st *int) int {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       XBrv     �               func Pop(var *int) int {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       XBrx     �               func Pop(var ...) int {5�_�                           ����                                                                                                                                                                                                                                                                                                                                         	       v       XBr�    �      	         	s := ""   	for s != "aaaaa" {   		fmt.Println("Value of s:", s)   		s = s + "a"   	}   		val := 1   
	Pop(&val)5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBr�    �   
            func Pop(var ... string) int {5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBt�     �      	         	5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                         
       v       XBt�     �      	         
	str := {}5�_�                       
    ����                                                                                                                                                                                                                                                                                                                                         
       v       XBt�     �      	         	str := {""}5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBt�     �      	         	str := {"ab"}5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBt�     �      	         	str := {"ab", ""}5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �      	         	str := {"ab", "dc"}5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �      	         	str := {"ab", "dc", ""}5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �      	         	str := {"ab", "dc", "jje"}5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �      	         	str := {"ab", "dc", "jje", ''}5�_�                           ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �      	         	str := {"ab", "dc", "jje", ""}5�_�                        "    ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �      	         #	str := {"ab", "dc", "jje", "jioe"}5�_�      !                  !    ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu	     �      
         	�      
       5�_�       "           !   	       ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �      
         	Pop()5�_�   !   #           "          ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu     �               	�             5�_�   "   $           #          ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu!     �               	for 5�_�   #   %           $          ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu-     �               		�             �               	for _, v := range var {}5�_�   $   &           %          ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu5     �               		fmt.Println()5�_�   %   '           &          ����                                                                                                                                                                                                                                                                                                                                         
       v       XBu6    �               		fmt.Println("")5�_�   &   +           '          ����                                                                                                                                                                                                                                                                                                                                         
       v       XBui    �      	         #	str := {"ab", "dc", "jje", "jioe"}5�_�   '   ,   (       +      
    ����                                                                                                                                                                                                                                                                                                                                                v       XBu�   
 �      	         %	str := []{"ab", "dc", "jje", "jioe"}5�_�   +   -           ,      	    ����                                                                                                                                                                                                                                                                                                                               	                 v       XBu�     �               func Pop(var ...string) int {5�_�   ,   .           -          ����                                                                                                                                                                                                                                                                                                                                                v       XBu�    �               	for _, v := range var {5�_�   -   /           .   	       ����                                                                                                                                                                                                                                                                                                                                                v       XBu�    �      
         		Pop(str)5�_�   .   0           /          ����                                                                                                                                                                                                                                                                                                                                                v       XBv�    �               func Pop(value ...string) int {5�_�   /   1           0   	       ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   	            	�   	          5�_�   0   2           1          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �                +	str := []string{"ab", "dc", "jje", "jioe"}5�_�   1   3           2          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �                	Pop(str...)5�_�   2   4           3           ����                                                                                                                                                                                                                                                                                                                                                  v        XB��     �   
               func Pop(value ... string) int {   	for _, v := range value {   		fmt.Println(v)   	}   		return 0   }5�_�   3   5           4           ����                                                                                                                                                                                                                                                                                                                                                  v        XB��     �      	         	�      	       5�_�   4   6           5      	    ����                                                                                                                                                                                                                                                                                                                                                  v        XB��     �      	         
	str := ""5�_�   5   7           6   	       ����                                                                                                                                                                                                                                                                                                                                                  v        XB�	     �      
         	strings.IndexFunc()5�_�   6   8           7   	       ����                                                                                                                                                                                                                                                                                                                                                  v        XB�     �      
         	strings.IndexFunc(str, )5�_�   7   9           8           ����                                                                                                                                                                                                                                                                                                                                                  v        XB�     �                �             5�_�   8   :           9          ����                                                                                                                                                                                                                                                                                                                                                  v        XB�/     �               func isAsc2()5�_�   9   ;           :          ����                                                                                                                                                                                                                                                                                                                                                  v        XB�2     �               func isAsc2(c int)5�_�   :   <           ;          ����                                                                                                                                                                                                                                                                                                                                                  v        XB�4     �               	�             �               func isAsc2(c int) bool {}5�_�   ;   =           <          ����                                                                                                                                                                                                                                                                                                                                                  v        XB�F     �               	if c < 2555�_�   <   >           =          ����                                                                                                                                                                                                                                                                                                                                                  v        XB�H     �               		�             �               	if c < 255 {}5�_�   =   ?           >           ����                                                                                                                                                                                                                                                                                                                                                  v        XB�P    �                5�_�   >   @           ?           ����                                                                                                                                                                                                                                                                                                                                                  v        XB�\    �                	"fmt"5�_�   ?   B           @           ����                                                                                                                                                                                                                                                                                                                                                  v        XB�^    �               	�             5�_�   @   C   A       B          ����                                                                                                                                                                                                                                                                                                                                                v   
    XB�x    �               	strings�             5�_�   B   D           C          ����                                                                                                                                                                                                                                                                                                                                         	       v   
    XB��     �               func isAsc2(c int) bool {5�_�   C   E           D          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �               func isAsc2(r int) bool {5�_�   D   F           E          ����                                                                                                                                                                                                                                                                                                                                                v       XB��    �               	if c < 255 {5�_�   E   G           F          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �      
         	�      
       5�_�   F   H           G   	       ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   	            		�   	          �               	for _, val := range str {}5�_�   G   I           H   
       ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   	            		strings.IndexFunc()5�_�   H   J           I   
       ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   	             		strings.IndexFunc(val, isAsc2)5�_�   I   K           J   
   #    ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   	            #		if strings.IndexFunc(val, isAsc2)5�_�   J   L           K   
   %    ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   
            			�   
          �   	            &		if strings.IndexFunc(val, isAsc2) {}5�_�   K   M           L          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   
            			fmt.Print()5�_�   L   N           M          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   
            			fmt.Printf()5�_�   M   O           N          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �   
            			fmt.Printf("")5�_�   N   P           O          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �               		�             5�_�   O   Q           P          ����                                                                                                                                                                                                                                                                                                                                                v       XB��     �               		fmt.Printf()5�_�   P   R           Q          ����                                                                                                                                                                                                                                                                                                                                                v       XB��    �                	strings.IndexFunc(str, isAsc2)5�_�   Q   S           R          ����                                                                                                                                                                                                                                                                                                                                                v       XB�#     �               	�             5�_�   R   T           S          ����                                                                                                                                                                                                                                                                                                                                                v       XB�$    �               	""5�_�   S   U           T   
       ����                                                                                                                                                                                                                                                                                                                                                v       XB�e     �   	   
          	for _, val := range str {5�_�   T   V           U          ����                                                                                                                                                                                                                                                                                                                                                v       XB�h     �                	}5�_�   U               V   
        ����                                                                                                                                                                                                                                                                                                                            
                      v        XB�i    �   	            %		if strings.IndexFunc(val, isAsc2) {   			fmt.Printf("?")   		}   		fmt.Printf(val)5�_�   @           B   A          ����                                                                                                                                                                                                                                                                                                                                         	       v   
    XB�t     �             �               
	st"rings"5�_�   '   )       +   (          ����                                                                                                                                                                                                                                                                                                                                                v       XBu�    �               func Pop(var ...int) int {5�_�   (   *           )          ����                                                                                                                                                                                                                                                                                                                                                v       XBu�     �      	         	str := []{}5�_�   )               *          ����                                                                                                                                                                                                                                                                                                                                                v       XBu�   	 �      	         	str := []{1, 2, 3, 4}5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             XAr�     �              5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             XAr�    �              5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �               		�               		for5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �                5�_�                             ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �               container_hero5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �               
	Pop(&val)5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �               		val := 15�_�          
                ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �                5�_�          	      
          ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �               
	Pop(&val)5�_�             
   	          ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �               		val := 15�_�             	             ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �                5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �               
	Pop(&val)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             XAt�     �             �               
	Pop(&val)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             XAtc     �               		�               		f5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             XAti     �               		fu5��