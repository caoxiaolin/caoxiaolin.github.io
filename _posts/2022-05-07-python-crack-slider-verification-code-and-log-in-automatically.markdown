---
layout:     post
title:      "python破解京东滑块验证码并自动登录"
tags:
    - python
    - slider verification code
    - opencv
    - pyautogui
---
### 滑块验证码

```python
import random
import time
import base64
import sys

import cv2
import numpy as np
import pyperclip
import sync
import pyautogui as pg
import win32gui
import win32clipboard as wc
from win32con import WM_INPUTLANGCHANGEREQUEST

def try_login(url='', account='', password=''):
    openMainPage(url)
    time.sleep(3)

    pyperclip.copy(account)

    pg.moveTo(935, 250, 1)
    pg.click()
    time.sleep(0.3)
    pg.hotkey('ctrl', 'a')
    pg.hotkey('ctrl', 'v')
    pg.click()
    time.sleep(0.3)

    pg.moveTo(925, 300, 1)
    pg.click()
    time.sleep(0.3)
    pg.write(password, 0.1)
    pg.click()
    time.sleep(0.3)

    pg.moveTo(1044, 346, 1)
    pg.click()

    time.sleep(3)
    move_huakuai()
    print(time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time())) + " 滑块破解成功")

def move_huakuai():
    while 1:
        print(time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time())) + " 开始破解滑块")
        pg.hotkey('f12')
        time.sleep(1)
        pg.moveTo(1289, 110, 0.2)
        pg.click()
        time.sleep(0.3)

        pg.moveTo(1050, 300, 0.5)
        pg.click()
        time.sleep(0.5)
        
        pg.hotkey('ctrl', 'shift', 'k')
        time.sleep(0.5)

        pg.write('clear()', 0.1)
        pg.hotkey('enter')

        time.sleep(0.1)
        pyperclip.copy('$(".JDJRV-bigimg img")[0].src')
        time.sleep(0.1)
        pg.hotkey('ctrl', 'v')
        pg.hotkey('enter')

        pg.moveTo(1311, 235, 0.5)
        pg.click()
        time.sleep(0.2)
        pg.click(button='right')
        time.sleep(0.1)
        pg.moveTo(1361, 285, 0.5)
        pg.click()
        time.sleep(0.6)
        img1 = get_copy_text()

        if not ";base64," in img1:
            pg.hotkey('f12')
            break

        pg.hotkey('ctrl', 'shift', 'k')
        time.sleep(0.5)

        pg.write('clear()', 0.1)
        pg.hotkey('enter')

        time.sleep(0.1)
        pyperclip.copy('$(".JDJRV-smallimg img")[0].src')
        time.sleep(0.1)
        pg.hotkey('ctrl', 'v')
        pg.hotkey('enter')

        pg.moveTo(1311, 235, 0.5)
        pg.click()
        time.sleep(0.2)
        pg.click(button='right')
        time.sleep(0.1)
        pg.moveTo(1361, 285, 0.5)
        pg.click()
        time.sleep(0.6)
        img2 = get_copy_text()

        save_huakuai("big.png", img1)
        save_huakuai("small.png", img2)

        #获取图片大小，主要是宽度，影响滑块移动距离
        pg.hotkey('ctrl', 'shift', 'k')
        time.sleep(0.5)

        pg.write('clear()', 0.1)
        pg.hotkey('enter')

        time.sleep(0.1)
        pyperclip.copy('$($(".JDJRV-bigimg")[0]).width()')
        time.sleep(0.1)
        pg.hotkey('ctrl', 'v')
        pg.hotkey('enter')

        pg.moveTo(1312, 214, 0.5)
        pg.doubleClick()
        time.sleep(0.1)
        pg.hotkey('ctrl', 'c')
        w = get_copy_text()
        time.sleep(1)
        
        pg.hotkey('ctrl', 'shift', 'e')
        pix = detect_displacement("small.png", "big.png")
        pix = pix * int(w) / 360  #缩放
        pg.moveTo(948+gr(3), 398+gr(10), 0.1)
        pg.drag(pix, 0, 1.5 + gr(1), pg.easeOutElastic)
        time.sleep(1)
        pg.hotkey('f12')
        time.sleep(10)


def save_huakuai(filename, b64data):
    b64data = b64data.replace('"', '')
    data = b64data.split(';base64,')[1]
    #print(data)
    with open(filename, "wb") as f:
        f.write(base64.b64decode(data))

def _tran_canny(image):
    """滑块处理，消除噪声"""
    image = cv2.GaussianBlur(image, (3, 3), 0)
    return cv2.Canny(image, 50, 150)


def detect_displacement(img_slider_path, image_background_path):
    """检测滑块位置，返回坐标"""
    # # 参数0是灰度模式
    image = cv2.imread(img_slider_path, 0)
    template = cv2.imread(image_background_path, 0)
    # 寻找最佳匹配
    res = cv2.matchTemplate(_tran_canny(image), _tran_canny(template), cv2.TM_CCOEFF_NORMED)
    # 最⼩值，最⼤值，并得到最⼩值, 最⼤值的索引
    min_val, max_val, min_loc, max_loc = cv2.minMaxLoc(res)
    top_left = max_loc[0]  # 横坐标
    return top_left

def get_copy_text(operate_open_close=False):
    suc = False
    count = 0
    while not suc:
        try:
            count = count + 1
            if not operate_open_close:
                wc.OpenClipboard()
            t = wc.GetClipboardData()
            #print('data')
            #print(t)
            if not operate_open_close:
                wc.CloseClipboard()
            return t
        except Exception as err:
            # If access is denied, that means that the clipboard is in use.
            # Keep trying until it's available.
            print(err)
            print(f'count:{count}')
            time.sleep(2)
            pass
            if err.winerror == 5:  # Access Denied
                # wait on clipboard because something else has it. we're waiting a
                # random amount of time before we try again so we don't collide again
                pass
            elif err.winerror == 1418:  # doesn't have board open
                pass
            elif err.winerror == 0:  # open failure
                pass
            else:
                print('ERROR in Clipboard section of readcomments: %s' % err)

                pass


def openMainPage(url=''):
    pg.moveTo(329, 62, 1)
    pg.hotkey('f5')
    time.sleep(6)
    pg.click()
    pyperclip.copy(url)
    pg.hotkey('ctrl', 'v')
    #pg.write(url, 0.1)
    pg.hotkey('enter')

def get_cookies():
    time.sleep(1)
    pg.hotkey('shift', 'f9')
    time.sleep(1)
    pg.moveTo(910, 338, 0.3)
    pg.click()
    time.sleep(0.4)
    pg.moveTo(1348, 286, 0.3)
    pg.click()
    time.sleep(0.5)
    left_x = 910
    right_x = 995
    y = 338
    step = 21
    count = 0
    s = ''
    time.sleep(1)
    while count < 19:
        pg.moveTo(left_x, y + count * step, 0.1)
        pg.doubleClick()
        time.sleep(0.1)
        pg.hotkey('ctrl', 'c')
        time.sleep(0.1)
        s += get_copy_text() + '='
        pg.moveTo(right_x, y + count * step, 0.1)
        pg.doubleClick()
        time.sleep(0.1)
        pg.hotkey('ctrl', 'c')
        time.sleep(0.1)
        s += get_copy_text() + '; '
        count += 1
        
    print(s)
    pg.hotkey('f12')
    return s

def main():
    pg.FAILSAFE = False
    try_login('https://thunder.jd.com/brandweb/brand/#/jdlszs', '账号', '密码')
    get_cookies()


if __name__ == '__main__':
    main()
```
