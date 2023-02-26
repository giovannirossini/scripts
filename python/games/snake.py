import pygame as pg
from random import randrange


window = 400
tile_size = 15
range = (tile_size // 2 + 20, window - tile_size // 2, tile_size)
screen = pg.display.set_mode((window, window))
pg.display.set_caption("Snake")
clock = pg.time.Clock()
pg.font.init()
get_random_position = lambda: (randrange(*range), randrange(*range))
time, snake_speed = 0, 110
snake_pos = get_random_position()
snake = pg.rect.Rect([0, 0, tile_size - 2, tile_size - 2])
snake.center = get_random_position()
length = 1
snake_body = [snake.copy()]
snake_color = 'white'
snake_dir = (0, 0)
food = snake.copy()
food.center = get_random_position()
food_color = 'red'
score = 0
score_font = pg.font.SysFont('arialunicode', 20)
score_color = 'white'
game_over = False
game_over_font = pg.font.SysFont('arialunicode', 40)
game_over_color = 'red'
game_over_text = game_over_font.render('Game Over', True, game_over_color)
game_over_text_rect = game_over_text.get_rect(center=(window/2, window/2))
self_collision = False

while not game_over:
    for event in pg.event.get():
        if event.type == pg.QUIT:
            game_over = True
        if event.type == pg.KEYDOWN:
            if (event.key == pg.K_UP or event.key == pg.K_w) and snake_dir != (0, tile_size):
                snake_dir = (0, -tile_size)
            if (event.key == pg.K_DOWN or event.key == pg.K_s) and snake_dir != (0, -tile_size):
                snake_dir = (0, tile_size)
            if (event.key == pg.K_LEFT or event.key == pg.K_a) and snake_dir != (tile_size, 0):
                snake_dir = (-tile_size, 0)
            if (event.key == pg.K_RIGHT or event.key == pg.K_d) and snake_dir != (-tile_size, 0):
                snake_dir = (tile_size, 0)

    screen.fill('black')
    pg.draw.rect(screen, food_color, food)
    [pg.draw.rect(screen, snake_color, body) for body in snake_body]
    if snake.center in [body.center for body in snake_body[:-1]]:
        self_collision = True

    if snake.left < 0 or snake.right > window or snake.top < 0 or snake.bottom > window or self_collision:
        score_text = score_font.render(f'Score: {score}', True, score_color)
        score_text_rect = score_text.get_rect(center=(window/2, window/2 + 33))
        screen.fill('black')
        screen.blit(score_text, score_text_rect)
        screen.blit(game_over_text, game_over_text_rect)
        if event.type == pg.KEYDOWN:
            if event.key == pg.K_RETURN:
                snake.center = get_random_position()
                snake_dir = (0, 0)
                time, snake_speed = 0, 110
                length = 1
                snake_body = [snake.copy()]
                score = 0
                self_collision = False
            if event.key == pg.K_ESCAPE:
                game_over = True
    if snake.center == food.center:
        food.center = get_random_position() if food.center not in [body.center for body in snake_body] else get_random_position()
        length += 1
        score += 1
        snake_speed = max(30, snake_speed - 3) if snake_speed > 60 else snake_speed - 1
        print(snake_speed)

    time_now = pg.time.get_ticks()
    if time_now - time > snake_speed:
        time = time_now
        snake.move_ip(snake_dir)
        snake_body.append(snake.copy())
        snake_body = snake_body[-length:]


    pg.display.flip()
    clock.tick(60)
